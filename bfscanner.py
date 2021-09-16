import os
import time


os.system ("apt-get install figlet")
os.system("sysctl -w vm.overcommit_memory=1")
os.system ("clear")

class style():          
    GREEN = '\033[32m'        
    WHITE = '\033[37m'

print(style.GREEN + "")

os.system ("figlet     BFScanner")   
print("---"*17)
print("		Go based port scanner")
print(style.WHITE + "")

target = input ("Target = ") 

t = open("target.txt","w")
t.write(target)
t.close()
time.sleep(1)
print("")

os.system("./scanner")

a=1
while a==1:
    time.sleep(1)
    if os.path.isfile("command.txt"):
        with open("command.txt") as file:
            command = file.read()
            file.close()
            
            os.system(command)
            print("\n\n-----------------------------------------------\n DONE, GOOD LUCK \n----------------------------------------------- ")
            os.remove("command.txt")
            if os.path.isfile("target.txt"):
                os.remove("target.txt")
            break
        

            

 

