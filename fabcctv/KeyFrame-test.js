const KeyFrame = require('./KeyFrame.js')
const VIDEO_FILE = './sample2.mp4';


var path = require('path');
const exif = require('exif-parser')
const fs = require('fs')

var ExifImage = require('exif').ExifImage;
let keyframe = new KeyFrame.KeyFrame();

function decode_base64(buf, filename) {
    // let buf = Buffer.from(base64str, 'base64');

    fs.writeFile(path.join(__dirname, '/public/', filename), buf, function(error) {
      if (error) {
        throw error;
      } else {
        console.log('File created from base64 string!');
        return true;
      }
    });
  }

keyframe.extractKey(VIDEO_FILE).then(keyframes => {
    // console.log(keyframes);
    console.log(keyframes[0])
    console.log('----------------')
    console.log(keyframes[0].image)
    // keyframes[0] = JSON.stringify(keyframes[0]);
    // keyframes[0] = JSON.parse(keyframes[0]);
    
    // decode_base64(keyframes[0].image, 'rane.jpg');
    
    
    // const buffer = fs.readFileSync(__dirname + '/public/rane3.jpg')
    // const parser = exif.create(buffer)
    // const result = parser.parse()
    
    // console.log(JSON.stringify(result, null, 2))

    try {
        new ExifImage({ image : path.join(__dirname, '/public/', 'rane3.jpg') }, function (error, exifData) {
            if (error)
                console.log('Error: '+error.message);
            else
                console.log(exifData); // Do something with your data!
        });
    } catch (error) {
        console.log('Error: ' + error.message);
    }


    
}).catch(err=>{
    console.log(err)
})





