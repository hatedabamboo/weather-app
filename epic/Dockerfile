FROM python:3.12-alpine

COPY . /app

WORKDIR /app

RUN pip install --no-cache-dir --break-system-packages -r requirements.txt

CMD ["python", "app.py"]
