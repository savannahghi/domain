// You can type/paste your script here

async function login({ graphql }) {
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
        "text": "terms will be available here soon",
        "flavour": "CONSUMER"
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
            "title": "Tips on how to keep yourself healthy",
            "body": "<h1>Content will be available here soon",
            "author": "Abiud Orina",
            "authorAvatar": "data:image/gif;base64,R0lGODlhAQABAIAAAP///wAAACH5BAEAAAAALAAAAAABAAEAAAICRAEAOw==",
            "heroImage": "data:image/gif;base64,R0lGODlhAQABAIAAAP///wAAACH5BAEAAAAALAAAAAABAAEAAAICRAEAOw==",
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
            "authorAvatar": "data:image/gif;base64,R0lGODlhAQABAIAAAP///wAAACH5BAEAAAAALAAAAAABAAEAAAICRAEAOw==",
            "heroImage": "data:image/gif;base64,R0lGODlhAQABAIAAAP///wAAACH5BAEAAAAALAAAAAABAAEAAAICRAEAOw==",
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
            "authorAvatar": "data:image/gif;base64,R0lGODlhAQABAIAAAP///wAAACH5BAEAAAAALAAAAAABAAEAAAICRAEAOw==",
            "heroImage": "data:image/gif;base64,R0lGODlhAQABAIAAAP///wAAACH5BAEAAAAALAAAAAABAAEAAAICRAEAOw==",
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
        "avatar": "1.3k"
    }, {
        "name": "Mental health",
        "avatar": "13k"
    }, {
        "name": "Ruaka self help",
        "avatar": "1.2k"
    }, {
        "name": "Ruiru",
        "avatar": "1.8k"
    }, {
        "name": "Kasarani",
        "avatar": "400"
    }, ]
}

self.addGraphQLResolvers({
    // Queries
    "Query.login": login,
    "Query.getCurrentTerms": getCurrentTerms,
    "Query.setUserPIN": genericBoolResponse,
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
    "Mutation.reviewTerms": genericBoolResponse,
    "Mutation.recordSecurityQuestionResponse": genericBoolResponse
})