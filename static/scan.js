const CameraAccess = Quagga.CameraAccess;

$(document).ready(function() {
  const el = $('#scanvideo')[0];

  Quagga.init({
    inputStream: {
      type: 'LiveStream',
      target: el,
      willReadFrequently: true,
    },
    decoder: {
      readers: ['code_128_reader'],
    },
    locator: {
      patchSize: 'x-large',
    },
  }, function(err) {
    if (err) {
      console.log(err);
      return;
    }
    CameraAccess.request(el, { facingMode: 'environment' }).then(function() {
      Quagga.start();
      console.log("Quagga2 Initialized");
    })
  });

  Quagga.onDetected(console.log);

});
