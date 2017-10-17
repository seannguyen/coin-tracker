FROM python:2.7 
ADD . /app
WORKDIR /app
RUN pip install -r requirements.txt
CMD ["python", "-m", "lib.pollers.balance"]