db = db.getSiblingDB('api-creator-documents');
db.createUser({
  user: 'user',
  pwd: 'password',
  roles: [
    { 
      role: 'readWrite', 
      db: 'api-creator-documents' 
    }
  ],
});

db.createCollection('test');
