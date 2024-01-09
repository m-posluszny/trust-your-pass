import numpy as np
from tensorflow.keras.preprocessing.sequence import pad_sequences
from keras.models import load_model
import pickle
from typing import Union
import json

strenght_mapper = [3,2,4,0,1]

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
    
def predict_password_strength(new_password):
    new_sequence = tokenizer.texts_to_sequences([new_password])
    new_padded_sequence = pad_sequences(new_sequence,232)
    prediction = model.predict(new_padded_sequence, verbose=None)
    predicted_class_index = np.argmax(prediction)
    predicted_strength = label_encoder.classes_[predicted_class_index]
    print(f' [process]  Password: "{new_password}", strenght: {predicted_strength}, mapped_strength:  {strenght_mapper[predicted_class_index]}' )
    result = { "_id":"fake_id", "strength":strenght_mapper[predicted_class_index] }
    json_str = json.dumps(result)
    return json_str
    
def main():	
    while True:            
        password = input('Enter password: ')
        result = predict_password_strength(password)
        print(f"JSON output: {result}",end="\n\n")


    
if __name__ == '__main__':
    try:
        print('To exit press CTRL+C');
        main()
    except KeyboardInterrupt:
        print(' Interrupted')
        
