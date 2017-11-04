#include "denarius.h"
#include <stdlib.h>
#include <stdint.h>
#include <string.h>
#include <stdio.h>

#include "sph_jh.h"
#include "sph_keccak.h"
#include "sph_echo.h"


void denarius_hash(const char* input, char* output)
{
    sph_jh512_context        ctx_jh;
    sph_keccak512_context    ctx_keccak;
    sph_echo512_context		ctx_echo;

    //these uint512 in the c++ source of the client are backed by an array of uint32
    //uint32_t hashA[16], hashB[16];
	
	uint8_t hash[64];

	sph_jh512_init(&ctx_jh);
	sph_jh512(&ctx_jh, input, 80); //inputlen
	sph_jh512_close(&ctx_jh, (void*) hash);

	sph_keccak512_init(&ctx_keccak);
	sph_keccak512(&ctx_keccak, (const void*) hash, 64);
	sph_keccak512_close(&ctx_keccak, (void*) hash);

	sph_echo512_init(&ctx_echo);
	sph_echo512(&ctx_echo, (const void*) hash, 64);
	sph_echo512_close(&ctx_echo, (void*) hash);

    memcpy(output, hash, 32);
	
}

