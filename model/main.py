import numpy as np
from tensorflow.keras.preprocessing.sequence import pad_sequences
from keras.models import load_model
from dotenv import load_dotenv
import pickle
from typing import Union
import json
import pika, sys, os

load_dotenv(dotenv_path="../tools/services.env")
strength_mapper = [3, 2, 4, 0, 1]

model_file__name = 'model.keras'
tokenizer_file_name = 'tokenizer.pickle'
encoder_file_name = 'label_encoder.pickle'
model = load_model(model_file__name)

# loading tokenizer
with open(tokenizer_file_name, 'rb') as handle:
    tokenizer = pickle.load(handle)

# loading label_encoder    
with open(encoder_file_name, 'rb') as handle:
    label_encoder = pickle.load(handle)


def predict_password_strength(jsonString):
    jsonDict = json.loads(jsonString)
    new_sequence = tokenizer.texts_to_sequences([jsonDict.get("password")])
    new_padded_sequence = pad_sequences(new_sequence, 232)
    prediction = model.predict(new_padded_sequence, verbose=None)
    predicted_class_index = np.argmax(prediction)
    predicted_strength = label_encoder.classes_[predicted_class_index]
    print(
        f' [process]  Password: "{jsonDict.get("password")}", strength: {predicted_strength}, mapped_strength:  {strength_mapper[predicted_class_index]}')

    print(f' [process]  Password: "{jsonDict.get("password")}"')
    return json.dumps({"_id": jsonDict.get("_id"), "strength": strength_mapper[predicted_class_index]})

def send_result(channel, modelOutQueueName, result):
    print(f" [x] Sent result: '{result}'")
    channel.basic_publish(exchange='', routing_key=modelOutQueueName, body=result)


def main():
    modelInQueueName = os.getenv("RABBITMQ_MODEL_IN_QUEUE_NAME")
    modelOutQueueName = os.getenv("RABBITMQ_MODEL_OUT_QUEUE_NAME")
    connection = pika.BlockingConnection(pika.URLParameters(os.getenv("RABBITMQ_CONNECTION_STRING")))
    channel = connection.channel()
    channel.queue_declare(queue=modelInQueueName, durable=True)
    channel.queue_declare(queue=modelOutQueueName, durable=True)

    def callback(ch, method, properties, body):
        decoded_string = body.decode("utf-8")
        print(f" [x] Received {decoded_string}")
        result = predict_password_strength(decoded_string)
        send_result(channel, modelOutQueueName, result)

    channel.basic_consume(queue=modelInQueueName, on_message_callback=callback, auto_ack=True)

    print(' [*] Waiting for messages. To exit press CTRL+C')
    channel.start_consuming()
    channel.connection.close()


if __name__ == '__main__':
    try:
        main()
    except KeyboardInterrupt:
        print('Interrupted')
        try:
            sys.exit(0)
        except SystemExit:
            os._exit(0)
