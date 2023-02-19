# nostrcheck
A very simple CLI tool to check the signature of a nostr note event

# Get it and build it
```sh
git clone git@github.com:vicariousdrama/nostrcheck.git
cd nostrcheck
go build
```

# Run it
3 samples provided
* sample0.json - nostr post by Mike Digler announcing Gossip 4.0
* sample1.json - nostr post by Vicarious Drama with ⚡️21⚡️ 
* sample2.json - same as sample1.json, but tampered to show signature check failure

```sh
./nostrcheck --notefile sample0.json
```
Expected Output
```
Using note posted by Mike Digler about Gossip as a sample
Formatted JSON of note event:
{
  "id": "8d600afa179801376a3c784bc79cec7753c768b8943911d3fb56fe29d63b221c",
  "pubkey": "ee11a5dff40c19a555f41fe42b48f00e618c91225622ae37b6c2bb67b76c4e49",
  "created_at": 1676437779,
  "kind": 1,
  "tags": [],
  "content": "Gossip 0.4.0 has been released!\n\nThis is the first release branch that will be supported with bug fixes.  I created packages for Debian and Windows and signed them with a PGP key that even I can't seem to verify (damn sub key nightmare!), so here is the SHA256SUMS.txt file signed under my nostr key:\n\na03207e4f357dc71e474b7cee6348031e5e3486c319a39f026b867ac50d29915  gossip_0.4.0_amd64.deb\n30ea3089e1c01341aeab703d41b335c12f3a906b120ee543f7af35f2602ab0d0  gossip.0.4.0.msi\n\nThis packaging/release process could use a lot of improvements.\n\nPersonally I still recommend pulling master and compiling.  I'll try to keep master on stable working commits.  But for those who like releases, well, there you go.\n\nCheers!",
  "sig": "32c4f489814608e1d4a8d8ba2ca18880cb96df216c883fd72aa80c4e8c1613a22b6138a69174481a5e9428fb25c19b9afb2eeecf9eff83ce58801918a7cd98d4"
}
Expected hash: 8d600afa179801376a3c784bc79cec7753c768b8943911d3fb56fe29d63b221c (matched)
Signature is valid from pubkey: ee11a5dff40c19a555f41fe42b48f00e618c91225622ae37b6c2bb67b76c4e49
```
