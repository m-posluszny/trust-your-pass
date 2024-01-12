FROM python:3.10

WORKDIR /app

COPY requirements.txt requirements.txt

RUN pip install --no-cache-dir -r /app/requirements.txt

COPY . .

CMD [ "python3", "main.py", "--env","./services.env" ]
