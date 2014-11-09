#!/usr/bin/python

import sys
import os
import re
import argparse

import tornado.ioloop
import tornado.web

import common.config as CONFIG
import common.redirectionHelper as REDIR


pathre = re.compile(r'^/([^/]+)/?(.*)$')

class AdminHandler(tornado.web.RequestHandler):
  def get(self):
    self.write(admin)

class GoHandler(tornado.web.RequestHandler):
  def get(self,path):
    alias = args = target = None
    
    m = pathre.search(path)
    if m:
      alias = m.group(1)
      args = m.group(2)
      if args == "":
        args = None
    else:
      self.send_error(status_code=406)
   
    try:
      target = store.get(alias)['target']
    except:
      self.send_error(status_code=204)

    try:
      redirection = REDIR.apply(target,args)
      self.redirect(redirection)
     
    except:
      self.send_error(status_code=500)
     
      
  def post(self,path):
    self.send_error(status_code=405)
  def put(self,path):
    self.send_error(status_code=405)
  def delete(self,path):
    self.send_error(status_code=405)


def main():
  parser = argparse.ArgumentParser(prog='GoSlash',description='URL abstraction service')
  parser.add_argument('-f',dest='configfile',metavar='config', nargs=1, help='Config file',required=True)
  parser.add_argument('--debug', dest='debug',nargs='?', type=bool,help='desc')
  
  args = parser.parse_args()
  
  if args.debug:
    DEBUG=True
  
  cfg = CONFIG.GoConfig()
  cfg.load(args.configfile[0])
  global store
  if cfg.store['engine'] == 'file':
    import common.storageFileHelper as STORAGE
    store = STORAGE.New(cfg.store['location'])
  #elif cfg.store['engine'] == 'redis':
  #elif cfg.store['engine'] == 'sql':
  #elif cfg.store['engine'] == 'mongodb':
  #elif cfg.store['engine'] == 'couchdb':



  application = tornado.web.Application([
    (r"(^(?!admin.*).*)$", GoHandler),
  ])
  application.listen(8888)
  admin = tornado.web.Application([
    (r"/admin(/.*)$", AdminHandler),
  ])
  admin.listen(8887)
  tornado.ioloop.IOLoop.instance().start()


if __name__ == "__main__":
	main()

