from time import sleep
import numpy as np
from tensorflow.keras.preprocessing.sequence import pad_sequences
from keras.models import load_model
from dotenv import load_dotenv
from pika.exceptions import AMQPConnectionError
import pickle
import json
import pika,  os
import argparse
import logging

logging.basicConfig(level=logging.INFO)
logging.info("Starting model service")
strength_mapper = [3, 2, 4, 0, 1]

model_file__name = 'model.keras'
tokenizer_file_name = 'tokenizer.pickle'
encoder_file_name = 'label_encoder.pickle'
model = None
tokenizer = None
label_encoder = None


def startup():
    print("Loading models")
    global model, tokenizer, label_encoder
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
    for i in range(9999):
        try:
            print("Connecting",os.getenv("RABBITMQ_CONNECTION_STRING"),i)
            connection = pika.BlockingConnection(pika.URLParameters(os.getenv("RABBITMQ_CONNECTION_STRING")))
        except AMQPConnectionError:
            sleep(1)
            continue
        break
    print("Connected")
    channel = connection.channel()
    channel.queue_declare(queue=modelInQueueName, durable=True)
    channel.queue_declare(queue=modelOutQueueName, durable=True)
    def callback(ch, method, properties, body):
        decoded_string = body.decode("utf-8")
        logging.info(" [x] Received %r" % decoded_string)
        result = predict_password_strength(decoded_string)
        send_result(channel, modelOutQueueName, result)

    channel.basic_consume(queue=modelInQueueName, on_message_callback=callback, auto_ack=True)

    startup() 
    print(' [*] Waiting for messages. To exit press CTRL+C')
    channel.start_consuming()
    connection.close()



if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument("--env", help="environment", default="./tools/services.env")
    args = parser.parse_args()
    print('loading env from: ' + args.env)
    load_dotenv(dotenv_path=args.env)
    main()
