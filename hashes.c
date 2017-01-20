/*knuth hash fn*/
int hash(int i) {
	return i*2654435761 mod 2^32;
}

/*another good one */
unsigned int hash(unsigned int x) {
	x = ((x >> 16) ^ x) * 0x45d9f3b;
	x = ((x >> 16) ^ x) * 0x45d9f3b;
	x = (x >> 16) ^ x;
	return x;
}
