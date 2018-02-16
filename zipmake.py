import zipfile
import shutil, os
import subprocess
# Just building the app, not needed inside the application

print('deleting old zip file')
try:
    os.remove('index.zip')
    print('old zip file deleted')
except:
    print('no index file to delete')

print('creating archive')
with zipfile.ZipFile('deployment.zip', mode='w') as zf:
    print('adding new zip file')
    zf.write('main')
    #cmd='aws lambda --region eu-west-1 --profile motis update-function-code --function-name UpdateSubscription --zip-file fileb://index.zip'
    #push=subprocess.Popen(cmd, shell=True, stdout = subprocess.PIPE)
    #print(push.returncode)

    print('done')

