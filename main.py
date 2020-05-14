
from colorama import Fore, Style
import glob
import os
import hashlib

valid_images = [".jpg", ".jpeg", ".png"]

def hashFromFile(f):
    md5 = hashlib.md5()
    BUF_SIZE = 65536  # lets read stuff in 64kb chunks!
    with open(f, 'rb') as file:
        while True:
            data = file.read(BUF_SIZE)
            if not data:
                break
            md5.update(data)
    return md5.hexdigest()

def getImages(path):
    image_list = []
    for ext in valid_images:
        for filename in glob.glob(path + '/*' + ext):  
            hashSum = hashFromFile(filename)
            if not hashSum in image_list:
                image_list.append(hashSum)
                print(Fore.GREEN,filename, hashSum, '\tok',Style.RESET_ALL)
            else:
                os.remove(filename)
                print(Fore.RED,'duplicada:\t', filename, hashSum, '\tdeletada',Style.RESET_ALL)


getImages(os.getcwd()+'/imagens')
