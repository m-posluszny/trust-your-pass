db = db.getSiblingDB('sample_db');

db.createCollection('passwords');

db.passwords.insertMany([
    {
        password: 'intel1',
        strength: 0,
        isProcessed: true
    },
    {
        password: 'elyass15@ajilent-ci',
        strength: 2,
        isProcessed: true
    },
    {
        password: 'hodygid757#$!23w',
        strength: 1,
        isProcessed: true
    },
    {
        password: 'notEstimatedYet',
        strength: -1,
        isProcessed: false
    }
]);