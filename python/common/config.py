import sys
import ConfigParser

DEBUG = True

def ConfigMap(config,section):
  map = {}
  options = config.options(section)
  for option in options:
    try:
      map[option] = config.get(section, option)
      if map[option] == -1:
        print("skip: %s" % option)
    except:
      print("exception on %s!" % option)
      map[option] = None
  return map

class GoConfig(object):
  def __init__(self):
    self.version = "1.0"
    self.format = "ini"
    self.store = None
    self.auth = None
    self.map = ConfigMap

  def load(self,cf):
    config = ConfigParser.ConfigParser()
    config.read(cf)
    try:
      self.store = self.map(config,"store")
    except ConfigParser.NoSectionError as e:
       print >> sys.stderr, "Config file has no [store] section configured !"
       exit(1)
    try:
      self.auth = self.map(config,"auth")
    except ConfigParser.NoSectionError as e:
       print >> sys.stderr, "Config file has no [auth] section configured !"
       exit(1)
    
