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
        "contentID": "some-content-id",
        "title": "Tips on how to keep yourself healthy",
        "body": "The coronavirus pandemic has affected our lives, our economy, and nearly every corner of the globe. Almost 4 billion vaccine doses have been administered worldwide; 53 for every 100 people. But the worldwide numbers of infections continue to rise, driven by the Delta variant with highly vaccinated regions like Western Europe and the United States, where cases are relatively low but climbing fast. As cases continue to surge, you can take some steps to keep yourself and your family safe. Here are some tips from our trusted science team.",
        "author": "Abiud Orina",
        "authorAvatar": "https://i.postimg.cc/9XpbrC25/profile-image.png",
        "heroImage": "https://i.postimg.cc/zvW46yxk/wellness.jpg",
        "contentType": "ARTICLE",
        "publicLink": "https://bewell.co.ke/",
        "isNew": true,
        "estimate": "4 minutes",
        "tags": ["Recommended", "Health", "Fitness"],
        "createdAt": "2021-08-23T06:42:05.085216Z",
    }, {
        "contentID": "some-content-id",
        "title": "What should my daily intake of calories be?",
        "body": "An ideal daily intake of calories varies depending on age, metabolism and levels of physical activity, among other things. Generally, the recommended daily calorie intake is 2,000 calories a day for women and 2,500 for men. The term calorie is commonly used as shorthand for kilocalorie. You will find this written as kcal on food packets. Kilojoules (kJ) are the equivalent of kilocalories within the International System of Units, and you'll see both kJ and kcal on nutrition labels. 4.2kJ is equivalent to approximately 1kcal.",
        "author": "Abiud Orina",
        "authorAvatar": "https://i.postimg.cc/9XpbrC25/profile-image.png",
        "heroImage": "https://i.postimg.cc/1t1XXvyz/lemon.jpg",
        "contentType": "ARTICLE",
        "publicLink": "https://bewell.co.ke/",
        "isNew": true,
        "estimate": "2 minutes",
        "tags": ["Recommended", "Health", "Fitness"],
        "createdAt": "2021-08-23T06:42:05.085216Z",
    }, ]
}

async function fetchContent({ graphql }) {
    return [{
            "contentID": "some-content-id",
            "title": "Tips on how to keep yourself healthy",
            "body": "The coronavirus pandemic has affected our lives, our economy, and nearly every corner of the globe. Almost 4 billion vaccine doses have been administered worldwide; 53 for every 100 people. But the worldwide numbers of infections continue to rise, driven by the Delta variant with highly vaccinated regions like Western Europe and the United States, where cases are relatively low but climbing fast. As cases continue to surge, you can take some steps to keep yourself and your family safe. Here are some tips from our trusted science team.",
            "author": "Abiud Orina",
            "authorAvatar": "https://i.postimg.cc/9XpbrC25/profile-image.png",
            "heroImage": "https://i.postimg.cc/zvW46yxk/wellness.jpg",
            "contentType": "ARTICLE",
            "publicLink": "https://bewell.co.ke/",
            "isNew": true,
            "estimate": "3 minutes",
            "tags": ["Recommended", "Health", "Fitness"],
            "createdAt": "2021-08-23T06:42:05.085216Z",
        }, {
            "contentID": "some-content-id",
            "title": "What should my daily intake of calories be?",
            "body": "An ideal daily intake of calories varies depending on age, metabolism and levels of physical activity, among other things. Generally, the recommended daily calorie intake is 2,000 calories a day for women and 2,500 for men.\n\n The term calorie is commonly used as shorthand for kilocalorie. You will find this written as kcal on food packets. Kilojoules (kJ) are the equivalent of kilocalories within the International System of Units, and you'll see both kJ and kcal on nutrition labels. 4.2kJ is equivalent to approximately 1kcal.",
            "author": "Abiud Orina",
            "authorAvatar": "https://i.postimg.cc/9XpbrC25/profile-image.png",
            "heroImage": "https://i.postimg.cc/T3wp4kdy/carrots.jpg",
            "contentType": "ARTICLE",
            "publicLink": "https://bewell.co.ke/",
            "isNew": false,
            "estimate": "3 minutes",
            "tags": ["Recommended", "Health", "Fitness"],
            "createdAt": "2021-08-23T06:42:05.085216Z",
        },
        {
            "contentID": "some-content-id",
            "title": "What is high blood pressure?",
            "body": "High blood pressure is often related to unhealthy lifestyle habits, such as smoking, drinking too much alcohol, being overweight and not exercising enough.\n\n Left untreated, high blood pressure can increase your risk of developing a number of serious long - term health conditions, such as coronary heart disease and kidney disease.",
            "author": "Abiud Orina",
            "authorAvatar": "https://i.postimg.cc/9XpbrC25/profile-image.png",
            "heroImage": "https://i.postimg.cc/tRKKf79F/blood-pressure.jpg",
            "contentType": "ARTICLE",
            "publicLink": "https://bewell.co.ke/",
            "isNew": false,
            "estimate": "2 minutes",
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
        "name": "Ruaka",
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
    "Query.fetchSuggestedGroups": fetchSuggestedGroups,

    /// Mutations
    "Mutation.setUserPIN": genericBoolResponse,
    "Mutation.acceptTerms": genericBoolResponse,
    "Mutation.recordSecurityQuestionResponse": recordSecurityQuestionResponses,
    "Mutation.saveContentItem": genericBoolResponse,
    "Mutation.likeContentItem": genericBoolResponse,
    "Mutation.shareContentItem": genericBoolResponse,
    "Mutation.setNickName": genericBoolResponse,
})