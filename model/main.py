import numpy as np
from tensorflow.keras.preprocessing.sequence import pad_sequences
from keras.models import load_model
import pickle
from typing import Union
import pika, sys, os

strenght_mapper = [3,2,4,0,1]
model_file__name = 'password_model_v2.keras'
tokenizer_file_name = 'tokenizer_v2.pickle'
encoder_file_name = 'label_encoder_v2.pickle'
model = load_model(model_file__name)
# loading
with open(tokenizer_file_name, 'rb') as handle:
    tokenizer = pickle.load(handle)
    
with open(encoder_file_name, 'rb') as handle:
    label_encoder = pickle.load(handle)
    
def predict_password_strength(new_password):
    new_sequence = tokenizer.texts_to_sequences([new_password])
    new_padded_sequence = pad_sequences(new_sequence,220)
    prediction = model.predict(new_padded_sequence, verbose=None)
    predicted_class_index = np.argmax(prediction)
    predicted_strength = label_encoder.classes_[predicted_class_index]
    result = f'Password: "{new_password}", strenght: {predicted_strength}, mapped_class:  [{strenght_mapper[predicted_class_index]}]' 
    print(result)
    return result
    
def send_result(result):
    connection = pika.BlockingConnection(pika.ConnectionParameters(host='localhost'))
    channel = connection.channel()
    channel.queue_declare(queue='result')
    channel.basic_publish(exchange='', routing_key='result', body=result)
    print(f" [x] Sent result: '{result}'")
    connection.close()

def main():
    connection = pika.BlockingConnection(pika.ConnectionParameters(host='localhost'))
    channel = connection.channel()

    channel.queue_declare(queue='password_to_process')

    def callback(ch, method, properties, body):
        decoded_string = body.decode("utf-8")
        print(f" [x] Received {decoded_string}")
        result = predict_password_strength(decoded_string)
        send_result(result)

    channel.basic_consume(queue='hello', on_message_callback=callback, auto_ack=True)

    print(' [*] Waiting for messages. To exit press CTRL+C')
    channel.start_consuming()
    
if __name__ == '__main__':
    try:
        main()
    except KeyboardInterrupt:
        print('Interrupted')
        try:
            sys.exit(0)
        except SystemExit:
            os._exit(0)
