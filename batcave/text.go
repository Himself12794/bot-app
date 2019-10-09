package batcave

const card = `
{
    "attachments": {
        "contentType": "application/vnd.microsoft.card.adaptive",
        "content": {
            "type": "AdaptiveCard",
            "version": "1.0",
            "body": [
                {
                    "type": "TextBlock",
                    "text": "Let's end world hunger, Be the Bridge",
                    "size": "large"
                },
                {
                    "type": "TextBlock",
                    "text": "*Sincerely yours,*"
                },
                {
                    "type": "TextBlock",
                    "text": "Bright Funds",
                    "separation": "none"
                },
                {
                    "type": "Input.ChoiceSet",
                    "id": "myCharity",
                    "style": "compact",
                    "isMultiSelect": false,
                    "value": "habitat",
                    "choices": [
                        {
                            "title": "Habitat for Humanity",
                            "value": "habitat"
                        },
                        {
                            "title": "Bike Ride for Bill",
                            "value": "bike"
                        },
                        {
                            "title": "Cars for Kids",
                            "value": "cars"
                        }
                    ]
                },
                {
                    "type": "Input.ChoiceSet",
                    "id": "mySelection",
                    "style": "expanded",
                    "isMultiSelect": true,
                    "value": "1",
                    "choices": [
                        {
                            "title": "Time",
                            "value": "1"
                        },
                        {
                            "title": "Money",
                            "value": "2"
                        },
                        {
                            "title": "Material",
                            "value": "3"
                        }
                    ]
                }
            ],
            "actions": [
                {
                    "type": "Action.Submit",
                    "title": "Search for Opportunities"
                }
            ]
        }
    },
    "roomId": "%s",
    "text": "text"
}`
