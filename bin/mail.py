#!/usr/bin/python
import argparse
import sys
#import smtplib
#from email.mime.text import MIMEText

from mailgun import Mailgun


def main():
	parser = argparse.ArgumentParser()
	parser.add_argument('--message', help='message',default='default')
	args = parser.parse_args(sys.argv[1:])
	message = args.message
	key = open('mailkey','r').read().strip()
	if key == '' or key is None:
		print('No Key')
	else:
		send(key, message)

def send(key, message):
	print(message)
	mailgun = Mailgun(apikey='key-2fba6c71e8d63bab7072d275297762c5', domain='mail.meirionconsulting.com')
	mailgun.send_message("Text, <gtd@mail.meirionconsulting.com>", "ian.miell@gmail.com", "GTD alert: " + message, message)
	##msg = MIMEText(message.encode('ascii','ignore'))
	#msg = MIMEText('test')
	#msg['Subject'] = "GTD alert: " + message
	#msg['From']    = "gtd@mail.meirionconsulting.com"
	#msg['To']      = "ian.miell@gmail.com"
	#s = smtplib.SMTP('smtp.mailgun.org', 587)
	#s.login('postmaster@mail.meirionconsulting.com', key)
	#s.sendmail(msg['From'], msg['To'], msg.as_string())
	#s.quit()
	print('email sent')

main()
