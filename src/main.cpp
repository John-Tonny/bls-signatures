#include <iostream>

#include "bls.hpp"


extern "C" {
#include "relic.h"
}


using namespace bls;

int main(void) 
{
  std::cout << "hello world!" << std::endl;


  vector<uint8_t> seed = {0,  50, 6,  244, 24,  199, 1,  25,  52,  88,  192,
                        19, 18, 12, 89,  6,   220, 18, 102, 58,  209, 82,
                        12, 62, 89, 110, 182, 9,   44, 20,  254, 22};


  PrivateKey sk = LegacySchemeMPL().KeyGen(seed);

  std::vector<uint8_t> sk1 = sk.Serialize(true);
  for(int i=0; i< sk1.size(); i++){
    std::cout << std::hex << static_cast<int>(sk1[i]) ;
  }
  std::cout << std::endl;

  G1Element pk = sk.GetG1Element();

  std::vector<uint8_t> pk1 = pk.Serialize(true);

  for(int i=0; i< pk1.size(); i++){
    std::cout << std::hex << static_cast<int>(pk1[i]) ;
  }
  std::cout << std::endl;


  // vector<uint8_t> message = {0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88,0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88};  // Message is passed in as a byte vector
  vector<uint8_t> message = {0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88,0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88};  // Message is passed in as a byte vecto

  G2Element signature = AugSchemeMPL().Sign(sk, message);

  // Verify the signature
  bool ok = AugSchemeMPL().Verify(pk, message, signature);

  std::cout << ok << std::endl;


  return 0;
}

