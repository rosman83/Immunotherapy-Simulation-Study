const superagent = require("superagent")

async function makeGetRequest(themethod, response) {
 await superagent
         .get( 'localhost:3000/probability' )
         .accept('application/json')
         .query({ method: themethod })
         .then(response => {
			  console.log(String(response.body.probdata))
			  console.log(String(response.body.moredata))
		})
}

for (i = 0; i < 10; i++) {
	makeGetRequest('ip')
}
