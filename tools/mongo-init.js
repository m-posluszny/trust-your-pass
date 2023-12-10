db = db.getSiblingDB('sample_db');

db.createCollection('passwords');

db.passwords.insertMany([
    {
        password: 'intel1',
        strength: 0
    },
    {
        password: 'elyass15@ajilent-ci',
        strength: 2
    },
    {
        password: 'hodygid757#$!23w',
        strength: 1
    },
    {
        password: 'notEstimatedYet',
        strength: -1
    }
]);