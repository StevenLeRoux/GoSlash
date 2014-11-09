import os
import csv

class New(object):
  def __init__(self,f):
    self.type="file"
    print "opening " + f
    store = csv.reader(open(f,"rb"))
    map = {}
    for row in store:
      if not map.has_key(row[0]):
        map[row[0]] = {'target': row[1],'user': row[2],'creationDate': row[3],'updateDate': row[4]}
    self.map = map
    print self.map

  def get(self,alias):
    return self.map[alias]
    
  
