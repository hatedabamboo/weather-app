FROM ubuntu:latest

RUN apt-get update && apt-get install -y python3 python3-pip

COPY . /app

WORKDIR /app

RUN pip install --break-system-packages -r requirements.txt

CMD ["python3", "app.py"]
