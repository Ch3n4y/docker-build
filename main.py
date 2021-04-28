import schedule, requests
import time, os

url = os.getenv('API')

def job():
    requests.get(url)


schedule.every(1).second.do(job)

while True:
    schedule.run_pending()
    time.sleep(1)
