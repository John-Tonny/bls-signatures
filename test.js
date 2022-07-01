
var loadBls = require("bls-signatures");

async function get(){
  var BLS = await loadBls();
  
  var seed = Uint8Array.from([
    0,  50, 6,  244, 24,  199, 1,  25,  52,  88,  192,
    19, 18, 12, 89,  6,   220, 18, 102, 58,  209, 82,
    12, 62, 89, 110, 182, 9,   44, 20,  254, 22
  ]);
  
  /*
  var sk = BLS.AugSchemeMPL.key_gen(seed);
  var pk = sk.get_g1();
  

  var message = Uint8Array.from([1,2,3,4,5]);
  var signature = BLS.AugSchemeMPL.sign(sk, message);
  
  let ok = BLS.AugSchemeMPL.verify(pk, message, signature);
  console.log(ok); // true
	
  var skBytes = sk.serialize();
  var publicKeyBytes = BLS.AugSchemeMPL.get_public_key(sk);
  var pkBytes = pk.serialize(false);
  var signatureBytes = signature.serialize(false);
  
  console.log(BLS.Util.hex_str(skBytes));
  console.log(BLS.Util.hex_str(publicKeyBytes));
  console.log(BLS.Util.hex_str(pkBytes));
  console.log(BLS.Util.hex_str(signatureBytes));
  */

  var sk1 = BLS.LegacySchemeMPL.key_gen(seed);
  var pk1 = sk1.get_g1();
  console.log(BLS.Util.hex_str(sk1.serialize()));
  console.log(BLS.Util.hex_str(BLS.LegacySchemeMPL.get_public_key(sk1)));

  var message1 = Buffer.from('1d314065b4ce9538cdffc5f4649e7068ccfc2cb0b45f5dd5df75d3bb31eed8b6', 'hex');
  var signature1 = BLS.LegacySchemeMPL.sign(sk1, message1);
  var signatureBytes1 = signature1.serialize(true);
  console.log(BLS.Util.hex_str(signatureBytes1));
  
  // var data = Buffer.from('908d5ee400058ae8869ac59313be26222838923bc666d5024bc5f0a6be284cb64e12fae2aecf026d6be8a70932d899de03f0e5e20449a3ada47cf0a65acee0c35915ef124519048996c74c19e4f6c08c50cf056461d171271988c9946548d79e', 'hex');
  var data = Buffer.from('377091f0e728463bc2da7d546c53b9f6b81df4a1cc1ab5bf29c5908b7151a32d', 'hex');
  var data1 = Uint8Array.from(data);
  console.log(data1);
  
  
  var sk2 = BLS.LegacySchemeMPL.key_gen(data1);
  console.log(BLS.Util.hex_str(sk2.serialize()));
  console.log(BLS.Util.hex_str(BLS.LegacySchemeMPL.get_public_key(sk2)));
  
  var message2 = Buffer.from('2d314065b4ce9538cdffc5f4649e7068ccfc2cb0b45f5dd5df75d3bb31eed8b6', 'hex');
  var signature2 = BLS.LegacySchemeMPL.sign(sk2, message2);


  // var signature2 = BLS.G2Element.from_bytes(data1);
  // console.log(signature2);
  let ok1 = BLS.LegacySchemeMPL.verify(pk1, message1, signature1);
  console.log(ok1); // true

  let ok2 = BLS.LegacySchemeMPL.verify(pk1, message1, signature2);
  console.log(ok2); // true

  sk1.delete();
  pk1.delete();
  signature1.delete();


}

get();
