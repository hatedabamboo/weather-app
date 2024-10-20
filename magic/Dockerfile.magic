FROM python:3.12-slim

COPY . /app

WORKDIR /app

RUN pip install --break-system-packages -r requirements.txt

CMD ["python", "app.py"]
