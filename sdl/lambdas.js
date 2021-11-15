// You can type/paste your script here

async function login({ graphql, args }) {
    return {
        "credentials": {
            "authCredentials": {
                "idToken": "some id token",
                "refreshToken": "some refresh token",
                "expiresIn": "3600",
            },
            "clientProfile": {
                "user": {
                    "userID": "some-user-id",
                    "userName": "Kowalski",
                    "displayName": "Kowalski",
                    "firstName": "Juha",
                    "lastName": "Kalulu",
                    "userType": "CLIENT",
                    "gender": "MALE",
                    "active": true,
                    "contacts": [{
                            "contactType": "PRIMARY",
                            "contact": "+254717356476",
                            "active": true,
                            "optedIn": true,
                        },
                        {
                            "contactType": "SECONDARY",
                            "contact": "+254712345678",
                            "active": true,
                            "optedIn": true,
                        },
                    ],
                    "languages": ["en", "sw"],
                    "termsAccepted": "true",
                },
                "treatmentEnrollmentDate": "21 Nov 2021",
                "clientType": "PMTCT",
                "active": true,
                "addresses": [{
                    "addressType": "POSTAL",
                    "text": "One Padmore",
                    "country": "Kenya",
                    "postalCode": "00300",
                    "active": true
                }],
                "facilityID": "some-facility-id",
                "clientCounselled": true
            },
            "pinChangeRequired": true
        },
        "code": "0",
        "message": "success",
    }
}

async function getCurrentTerms({ graphql }) {
    return {
        "termsID": "some-terms-id",
        "text": "Terms will be available here soon",
    }
}

async function genericBoolResponse({ graphql }) {
    return true
}


async function getSecurityQuestions({ graphql }) {
    return [{
            "questionStem": "What are the last 4 digits of your CCC number?",
            "description": "Please provide the last 4 digits of your clinic number",
            "flavour": "CONSUMER",
            "active": true,
            "responseType": "NUMBER",
        },
        {
            "questionStem": "Which month did you start your treatment?",
            "description": "Enter the month you started your treatment",
            "flavour": "CONSUMER",
            "active": true,
            "responseType": "DATE",
        },
        {
            "questionStem": "Which county is your clinic located?",
            "description": "Enter the name of the county in small letters",
            "flavour": "CONSUMER",
            "active": true,
            "responseType": "STRING",
        },
    ]
}

async function sendOTP({ graphql }) {
    return "111222"
}


async function verifyOTP({ graphql }) {
    return {
        "response": true,
        "code": null
    }
}

async function fetchRecentContent({ graphql }) {
    return [{
            "title": "5 proven ways to prevent COVID-19:Our experts reveal",
            "body": "<h1>Content will be available here soon",
            "author": "Abiud Orina",
            "authorAvatar": "https://i.postimg.cc/9XpbrC25/profile-image.png",
            "contentType": "ARTICLE",
            "publicLink": "https://bewell.co.ke/",
            "isNew": true,
            "estimate": "5 minutes",
            "tags": ["Recommended", "Health", "Fitness"],
            "createdAt": "2021-08-23T06:42:05.085216Z",
        }, {
            "title": "Tips for better and healthy living",
            "body": "<h1>Content will be available here soon",
            "author": "Abiud Orina",
            "authorAvatar": "https://i.postimg.cc/9XpbrC25/profile-image.png",
            "contentType": "ARTICLE",
            "publicLink": "https://bewell.co.ke/",
            "isNew": true,
            "estimate": "5 minutes",
            "tags": ["Recommended", "Health", "Fitness"],
            "createdAt": "2021-08-23T06:42:05.085216Z",
        },
        {
            "title": "How to keep fit",
            "body": "<h1>Content will be available here soon",
            "author": "Abiud Orina",
            "authorAvatar": "https://i.postimg.cc/9XpbrC25/profile-image.png",
            "contentType": "ARTICLE",
            "publicLink": "https://bewell.co.ke/",
            "isNew": true,
            "estimate": "5 minutes",
            "tags": ["Recommended", "Health", "Fitness"],
            "createdAt": "2021-08-23T06:42:05.085216Z",
        },
    ]
}

async function fetchContent({ graphql }) {
    return [{
            "title": "Tips on how to keep yourself healthy",
            "body": "<h1>Content will be available here soon",
            "author": "Abiud Orina",
            "authorAvatar": "https://i.postimg.cc/9XpbrC25/profile-image.png",
            "contentType": "ARTICLE",
            "publicLink": "https://bewell.co.ke/",
            "isNew": true,
            "estimate": "5 minutes",
            "tags": ["Recommended", "Health", "Fitness"],
            "createdAt": "2021-08-23T06:42:05.085216Z",
        }, {
            "title": "Tips for better and healthy living",
            "body": "<h1>Content will be available here soon",
            "author": "Abiud Orina",
            "authorAvatar": "https://i.postimg.cc/9XpbrC25/profile-image.png",
            "contentType": "ARTICLE",
            "publicLink": "https://bewell.co.ke/",
            "isNew": true,
            "estimate": "5 minutes",
            "tags": ["Recommended", "Health", "Fitness"],
            "createdAt": "2021-08-23T06:42:05.085216Z",
        },
        {
            "title": "How to keep fit",
            "body": "<h1>Content will be available here soon",
            "author": "Abiud Orina",
            "authorAvatar": "https://i.postimg.cc/9XpbrC25/profile-image.png",
            "contentType": "ARTICLE",
            "publicLink": "https://bewell.co.ke/",
            "isNew": true,
            "estimate": "5 minutes",
            "tags": ["Recommended", "Health", "Fitness"],
            "createdAt": "2021-08-23T06:42:05.085216Z",
        },
    ]
}


async function fetchSuggestedGroups({ graphql }) {
    return [{
        "name": "Ruaraka",
        "avatar": "https://i.postimg.cc/9XpbrC25/profile-image.png",
        "members": "20"
    }, {
        "name": "Mental health",
        "avatar": "https://i.postimg.cc/9XpbrC25/profile-image.png",
        "members": "50"
    }, {
        "name": "Ruaka self help",
        "avatar": "https://i.postimg.cc/9XpbrC25/profile-image.png",
        "members": "210"
    }, {
        "name": "Ruiru",
        "avatar": "https://i.postimg.cc/9XpbrC25/profile-image.png",
        "members": "10"
    }, {
        "name": "Kasarani",
        "avatar": "https://i.postimg.cc/9XpbrC25/profile-image.png",
        "members": "15"
    }, ]
}

async function recordSecurityQuestionResponses({ graphql }) {
    return [{
        "securityQuestionID": "1",
        "isCorrect": true,
    }, {
        "securityQuestionID": "2",
        "isCorrect": true,
    }, {
        "securityQuestionID": "3",
        "isCorrect": true,
    }, ]
}

self.addGraphQLResolvers({
    /// Queries
    "Query.login": login,
    "Query.getCurrentTerms": getCurrentTerms,
    "Query.getSecurityQuestions": getSecurityQuestions,
    "Query.sendOTP": sendOTP,
    "Query.resendOTP": sendOTP,
    "Query.verifyOTP": verifyOTP,
    "Query.fetchRecentContent": fetchRecentContent,
    "Query.fetchContent": fetchContent,
    "Query.setNickName": genericBoolResponse,
    "Query.fetchSuggestedGroups": fetchSuggestedGroups,

    /// Mutations
    "Mutation.setUserPIN": genericBoolResponse,
    "Mutation.acceptTerms": genericBoolResponse,
    "Mutation.recordSecurityQuestionResponse": recordSecurityQuestionResponses,
    "Mutation.saveContentItem": genericBoolResponse,
    "Mutation.likeContentItem": genericBoolResponse,
    "Mutation.shareContentItem": genericBoolResponse,
})