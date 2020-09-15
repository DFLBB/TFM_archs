import sys
import os
from keras.models import load_model
import urllib.request
from PIL import Image
from keras.models import load_model
import urllib.request
from PIL import Image
import pandas as pd
import numpy as np
from numpy import expand_dims
import os
import sys
import re
import shutil
import keras
from tensorflow.keras.layers import Dense,GlobalAveragePooling2D
from tensorflow.keras.applications import MobileNet, VGG16
from tensorflow.keras.models import Model
from tensorflow.keras.preprocessing.image import ImageDataGenerator
from tensorflow.keras.applications.mobilenet import preprocess_input as preprocess_mobilenet
from tensorflow.keras.applications.densenet import preprocess_input as preprocess_densenet
from tensorflow.keras.applications.nasnet import preprocess_input as preprocess_nasnet
from tensorflow.keras.applications.vgg16 import preprocess_input as preprocess_vgg16
from tensorflow.keras.applications.inception_resnet_v2 import preprocess_input as preprocess_inceptionResNetV2
from tensorflow.keras.optimizers import Adam
from sklearn.model_selection import train_test_split
from sklearn.metrics import classification_report
import matplotlib.pyplot as plt
import logging
from random import randrange
import pydot
from keras.utils.vis_utils import plot_model
from IPython.display import Image
from IPython.core.display import HTML
import collections
import math
import os
import time
import tensorflow as fl
from tensorflow.keras.preprocessing import image
from tensorflow.keras.applications import VGG16
from tensorflow.keras.layers import Dense, Conv2D, MaxPooling2D, Flatten, Dropout, GlobalAveragePooling2D
from tensorflow.keras.models import Model, Sequential
from PIL import Image

model = keras.models.load_model('/home/ec2-user/Notebooks/modelMobileNetV2.h5')
path = sys.argv[1]
print("El path es: ", path)
images = Image.open(urllib.request.urlopen(path))
images

def predict_image(path):
    img = path.resize((224, 224))
    data = expand_dims(path, 0)
    data = preprocess_mobilenet(data)
    preds = model.predict(data)
    pred = np.argmax(preds)
    return pred

indice = predict_image(images)

#Obtenemos las clases
classes = []
count = 0
data = []

datasetPath = "/home/ec2-user/images/Images"
logging.info(" Buscando la informacion del dataset")
for root, dirnames, filenames in os.walk(datasetPath):
    if(not classes):
        classes = dirnames
        continue        
    data.append((classes[count],filenames))
    count+=1  
    
print("raza:"+ classes[indice])