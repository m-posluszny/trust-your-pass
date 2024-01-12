db = db.getSiblingDB('sample_db');

db.createCollection('passwords');

db.passwords.insertMany([
    {
        password: 'elyass15@ajilent-ci',
        strength: 2,
        isProcessing: false
    },
    {
        password: 'hodygid757#$!23w',
        strength: 1,
        isProcessing: false
    }
]);