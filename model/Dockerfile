FROM python:3.10

WORKDIR /app

COPY requirements.txt requirements.txt

RUN pip install -r /app/requirements.txt

COPY . .

CMD [ "python3", "main.py", "--env","./services.env" ]
