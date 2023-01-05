import smtplib

server = smtplib.SMTP('smtp.gmail.com',587)

server.starttls()

server.login('sivalovely1243@gmail.com','lovely143')

server.sendmail('sivalovely1243@gmail.com', 'sivalovely1243@gmail.com', 'Mail sent from python')
print('mail sent')