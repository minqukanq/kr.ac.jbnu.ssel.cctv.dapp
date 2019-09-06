const KeyFrame = require('./KeyFrame.js')
const VIDEO_FILE = './sample2.mp4';

var path = require('path');
const exif = require('exif-parser')
const fs = require('fs')

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

    // decode_base64(keyframes[0].image, 'filename.jpg');
    
}).catch(err=>{
    console.log(err)
})





