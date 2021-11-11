// You can type/paste your script here

async function login({ graphql }) {
    return {
        "credentials": {
            "idToken": "some id token",
            "refreshToken": "some refresh token",
            "expiresIn": "3600",
            "user": {
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
            }
        },
        "code": "0",
        "message": "success",
    }
}

async function listTerms({ graphql }) {
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
        },
        {
            "questionStem": "Which month did you start your treatment?",
            "description": "Enter the month you started your treatment",
            "flavour": "CONSUMER",
            "active": true,
        },
    ]
}

async function sendOTP({ graphql }) {
    return "111222"
}

self.addGraphQLResolvers({
    // Queries
    "Query.login": login,
    "Query.listTerms": listTerms,
    "Query.setUserPIN": genericBoolResponse,
    "Query.getSecurityQuestions": getSecurityQuestions,
    "Query.sendOTP": sendOTP,
    "Query.resendOTP": sendOTP,


    /// Mutations
    "Mutation.setUserPIN": genericBoolResponse,
    "Mutation.reviewTerms": genericBoolResponse,
    "Mutation.recordSecurityQuestionResponse": genericBoolResponse
})