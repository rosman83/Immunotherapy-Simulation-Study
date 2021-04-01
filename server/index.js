var express =  require('express');
var weighted = require('weighted')

const app = express();

app.use(express.json());
const port = 3000



app.get('/probability', (req, res) => {
    // Identify what probability the machine model needs to grab
    console.log("Request Recieved")
    let method = req.body.method
    if ( method == 'ip' ) {
        let immunooptions = ['adverse', 'none']
         , immunoweights = [0.99, 0.01]
    
        let calcValues = String(weighted.select(immunooptions, immunoweights))
    } else if ( method == 'cp' ) {
        let chemooptions = ['adverses', 'nones']
        , chemoweights = [0.01, 0.99]

        let calcValues = String(weighted.select(chemooptions, chemoweights))
    } else if ( method == 'fip' ) {
        let fimmunooptions = ['fatality', 'none']
        , fimmunoweights = [87, 9913]

        let calcValues = String(weighted.select(fimmunooptions, fimmunoweights))
    } else if ( method == 'fcp' ) {
        let fchemooptions = ['fatalitys', 'nones']
        , fchemoweights = [128, 9872]

        let calcValues = String(weighted.select(fchemooptions, fchemoweights))
    }
    return res.send(calcValues);
});

   
  // Log that the application is running successfully
  app.listen(port, () =>
    console.log(`Server Side handler for Immunotherapy Machine Model is running at Port: ${port}!`),
  );

