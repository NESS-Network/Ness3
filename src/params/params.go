package params

/*
CODE GENERATED AUTOMATICALLY WITH FIBER COIN CREATOR
AVOID EDITING THIS MANUALLY
*/

const (
	// Distribution locking parameteres

	// MaxCoinSupply is the maximum supply of coins
	MaxCoinSupply uint64 = 200000000
	// DistributionAddressesTotal is the number of distribution addresses
	DistributionAddressesTotal uint64 = 100
	// DistributionAddressInitialBalance is the initial balance of each distribution address
	DistributionAddressInitialBalance uint64 = MaxCoinSupply / DistributionAddressesTotal
	// InitialUnlockedCount is the initial number of unlocked addresses
	InitialUnlockedCount uint64 = 50
	// UnlockAddressRate is the number of addresses to unlock per unlock time interval
	UnlockAddressRate uint64 = 5
	// UnlockTimeInterval is the distribution address unlock time interval, measured in seconds
	// Once the InitialUnlockedCount is exhausted,
	// UnlockAddressRate addresses will be unlocked per UnlockTimeInterval
	UnlockTimeInterval uint64 = 31536000 // in seconds
)

var (
	// Transaction verification parameters

	// UserVerifyTxn transaction verification parameters for user-created transactions
	UserVerifyTxn = VerifyTxn{
		// BurnFactor can be overriden with `USER_BURN_FACTOR` env var
		BurnFactor: 10,
		// MaxTransactionSize can be overriden with `USER_MAX_TXN_SIZE` env var
		MaxTransactionSize: 32768, // in bytes
		// MaxDropletPrecision can be overriden with `USER_MAX_DECIMALS` env var
		MaxDropletPrecision: 3,
	}
)

// distributionAddresses are addresses that received coins from the genesis address in the first block,
// used to calculate current and max supply and do distribution timelocking
var distributionAddresses = [DistributionAddressesTotal]string{
	"2BmHcwsZGfsBujFMzu56dU7gX6PGeZcLY33",
	"CCBkWFPekxmxUGLy81UMkNWBKam2kD2rra",
	"sTMg5Cvj7TZuuJ32P6o4apcjBdgf5mpQmu",
	"tkfrmtVoPA7sCXAqP1LLqLcdFMje6YQTsb",
	"MdY5ubZauLoZwwuZmBUUzMnDoyean3E2dg",
	"L91j5V5jqWEGmRx8VmN8SWMdgqHv5mkJtJ",
	"vkCBJKsTESppRWaQBZ78ZTsRtPJZWpd1YY",
	"RvqstBznCNfskWPFu9FLiGVRHNsLQQmXNR",
	"2ajyDvDAvMaz47qL2qzpMctFQjLfy1xdDaG",
	"zZpg4MgcHqthGXcdikPA3rEWo79MtZYuw1",
	"2gLVu53rv2CcCWaSbdC5fUbSgjnsMTx4nmg",
	"QccYggYfapHpaWaUTw1Hmj4bCgLMgSdp8K",
	"2MwnXCtye1J1XpkfErVFW2P9b1zrGGwdrZj",
	"Z3ueeeZTvvyX2BkygS6zeWXVR9Bc4PzQbG",
	"2UF9CwxeckPwp34DvvtYt4VjbLqAo6S35Th",
	"XimbQVNKG7AQUzLQMzFRUUGmgfpvBs8oCr",
	"XKkmDU3CyWs9FteAL2qiHpYeZhsoCVGC3h",
	"a7SHwgN4sN2unFvaXyigzffunCsBFGzEJW",
	"2B4AFQQAFPa7nekeJHC1VUkTXZ41Pcjg8vz",
	"hGhN4ptJtwtonSE8eMo7zphQoTmBvzzoCw",
	"YKiBqnM5WTxi7JtDnQHKchBJ794sAERszZ",
	"pJNx9fHjfkzvJbtACotx1dwQiN79Ad13pi",
	"2CBbxbVK7KAnUk1KnSogEWgMzmZbYmAepdM",
	"2A2bTC8aKssa4Edg2xTNnYhM4VSXr71xjmx",
	"Myxf1YNqRU9fcJcpLxC15PcsNqcCBi55Sf",
	"y8oAVTHHEte5tqJBSoEBqKN82nnDEZQFwW",
	"2WFbjRbppM5fvWT7WxytLw9hq87fxAvEiqx",
	"UMgZdHN4Sc25zzE55wTn2ZY432YG8edoos",
	"pw3UG6p653MS6z1QJiTPfgCyFskiDfSyEQ",
	"kayUntSjGYoTzmXtRaXs89gigzvBGdV5hf",
	"KicgPKCAEoaoj5yKCeEmFaRKRqA91kzxn2",
	"23mL9WEBSpuUuDPYmpg4dUWfs31hJzsU9VD",
	"21Kw2EVLzipd33M6D11SJRAtQjwQJbP7fF2",
	"pf4YRuKMGHkakZe8QmjZk8tV2EUUazDhgU",
	"2XLp3ZPEbPrtK2tQhtwzggzNVhZYgZbXG6M",
	"jWxFUVngXz4B9XTKPRMaqFuQazeTgnGB1K",
	"PEhXGmmu2DXnqZKRGLiaLUhDJgPWArU8Ts",
	"LFkbb8aKskewynbokDf1QyB2ugW5CwEr7U",
	"UtNb7jo9uUPJDYoHZvMmpo35rNQLEoZuR5",
	"23np7RehaUojs1bdLNfSAqum8jnuXvvRBsC",
	"YaDEPmNyndCGSPZJ7h5uh5WMEZ6Bm2JY8M",
	"2NmXsuTTcev6iyWxBWFrsG6rmCAj7Ree9Pb",
	"xi8jgGUwEn8qq1b5Y5xhJKRL3PebBGXAjH",
	"rkKQRiE7WTaULVoQQcjhpVL34YJ3Jwday1",
	"RQHWPCkNmHQ6r54hkKy5minw5ca5dJeBqt",
	"BVpUUfuUjvFJybwgDKNP3Y27s1eVzdS5Mm",
	"7C26Jp1PgkF8dWqzqQRMRHGCMafkTqpy71",
	"Cvc3iPofoJm7BDH8kCREcWbs8288goia5W",
	"2CG71n3LfiworsfW6SbgMaez59vb5UvyimH",
	"NGPnu4715Ca6T8VET7AujN7tmJGeqSUnoB",
	"hgfHtRvkg3dmEG4ZzTo2NvSkdPLB4BJZ76",
	"L5PTHpDQnLm5dtu6cqiaXfbUgTqEMs29gU",
	"2WwRU2nEkz2Bc3NqUV1SC71hC2g1EHnZ4CU",
	"ScF1pkDKYjzBJCNiL5GXdXRM5AqkS7xQYe",
	"i8sNBhqmzi9UfgggpeyZa9rM8QUyq6KG3k",
	"ATC5MCLg8sYNVvJ7VrbPTULkk8CsNNr99y",
	"2KAnXMppK5Y56H7ocpKZ62BPJQp3T1oLd9S",
	"qBAqzbZ4UjN47QGQoRdfvVdauju1BHghrW",
	"284yHHeZVBjJVXZTpeXT5UpbAQzrDVRYLp1",
	"uwGfb6jZo9epGHWi2GMPcggXX69AmDmjCt",
	"NnodkW7L9WUEPQZrp7EjcVibdRT9JqV8ie",
	"4mSndWxXTsY4DfMdxKaTGegxUA8YK2Ri8J",
	"255RnXwvXzchTa3WTh6X9tyzEPxQRrBaN7E",
	"xPGvitaW8ogbwTurP6duAkkrtokAempMw1",
	"2RN7T8ojbNXF1panqTy6e6EbdwFFcC4fvR",
	"2ULACq4hp2Fwq7WFo28UawqiKBDwMQmAK8N",
	"K8ZMKYKdRvNcNVWGiUvaNn8GMesV2PX6rS",
	"WfxJVFJspqeti16Eyj6u6oTXL18E5HVDCt",
	"2SrERQ482pHjNsyCnuV88zJd9xdCrN3z3Ky",
	"UsacH2e1sSamtQJHMzyLpcYB6RYLXEZgrN",
	"2eAeAiyaGfaVf6tVFJa9aSHHxrvJfDnfhhT",
	"2RNkdz7QLT4Mp1EWrF8cfNTQvjLyKjo4Ndz",
	"2PtL4VdQ3JyQWaA114nAc3etj5h3y9a45u3",
	"rrRYTpBtATmJpJTZqzsh699gYuwKPxMV9S",
	"rcNmS5YCPxDMhQiizMYnBrjowioHKm8H9p",
	"2T9NTVmW9f9ueHThKcgyUpK4m8omfqmNe7d",
	"2GyqmNdTcsa4b8y6AFBvYNw44KP7ofiCkA7",
	"2Ca6DwXWXhcfwDY1STvJPvfiKrGreTevRra",
	"tVGtnjrFWSof2LbHh5XZtFomNQf9cXmCdo",
	"2mTGN5xdAU8GZdM5QJTmppc7JEbeEb2Lnkv",
	"2aSMkGeHykXidku6SY1236JbobPwQ5YMkzJ",
	"BoX6TBfKCWuMtwVQpxKajkhcZh2WQsqYWK",
	"wY8juxKi59KaMgquUw81KHLDqiHYjoK9sC",
	"Eu1sa3P3WyjPzhnsioFQRMRyRAmrVGQdz6",
	"gikXmbrTLZbuBpdzwxhoaMNiUt3YZrywua",
	"2J4GEt3FATZBvN8wGFxXjfuyvyeHwowc6jd",
	"2WBBuPSvtfTP3mofk4k77chbLFJCnM7LtCM",
	"uH8a6fsCqnKjniN5vCAbKieUCZrHyEAxkH",
	"2m3Eae9eXAaKdXkXs9q9FcZyYTY1vezUegR",
	"XWaGjQjaFufTi7K8hibpgmVa81PUag5Msf",
	"28KFLiigteRmUdBPBLcEcvkcg8GKDs1nbxJ",
	"21aR2yTqDwHdXUQekxzCR2LQ2dzjLfN8TuF",
	"25KEti7VLVFD7f6CB14KLq955QyuJ76HBf8",
	"2cWAwhjdZWPJbywcuv1yuFyvw89vSnwCokQ",
	"2Tpe1sJw6wmSxF7vHy2j4Jvne96pWfDpR5a",
	"2bJ98k7u5D5gvJudebuJZCC9XABFRUcDeFk",
	"C5Eziwzhkw1Up3pZX3JnWhAkLSgTnP7WEr",
	"2BU3aDFgNSRdJHdzEPZbQSRMxcxqKvzd8M8",
	"2AAsM7oWw9p6RtnMPtpSAY5FeCBcexyvchC",
	"41LnUTdrknNm8J1fdohS7au97C8QVxNETD",
	"pntJVpzCeFgsYpDjpFLYWJnKkP4SmUCEmP",
	"28DkRyqhB4KQVzC2MvBnTzoA7CtxbUkw6xq",
	"AoUydYq9AzbfyANxPzgRkRK7HutfgzYXmh",
	"2kgEYrUXmxWa1GKM5D9VC4bayEuqS3A8kBX",
	"2CpxQz5GKNtpbZiYV3YmGirt4F5FyGXS3c2",
	"7ox2NCeMUBx2G5LVcWJuGKjjXotkoFAGSL",
	"NcUXBmQAq97af9ZWnPD95n9rhgQTp559vH",
	"UxgbSq2TNxhZSdNSVLGUq9p5UcPCwwWmiL",
	"2HyuxaSsToQM58W8uFuamKqJsM4jBtX39PK",
	"vaAoDPMAp7hsNroUKgqp6YAWkMQhy5ZEKz",
	"2fHAw5AeyeFhKDenm4MCgwiRyxM1dvEc7fK",
	"cGq7By7xE1cK1A1cRDkfdCz99J3tdc1sxa",
	"vCT6ChxBXNb5P6xPbQPVQwQ4bV2GFK1o8G",
	"2BKqXzkSEFuiyGaXP2E4zePFwZdemD7oq6z",
	"iLmw9PLCtoA1PXtKyKTtn9btUGNASLAwjc",
	"2dtXhdUPsHmMpHNz4GmBauMi4kde2G9R4jY",
	"LyQSnX8R3wNm17c9FTRnLcby1wZXiU2NhJ",
	"26vKc3rdiw12ybp3NnMBigoYE7oNte3Y3D",
	"sgJQ2qgMQYXAV7ULBTpfhnsoNRMk9cBUnQ",
	"2RqmB4t86WU49M5V7W4ZfFhmHbjLHrKn5Gp",
	"oWZQepcw6gXbue1B1QoqambyPthgRocz4C",
	"8qDuuNBhBfHfx9gE5A14L5DjSP4e8VF46u",
	"Ke8hi4wbtGUG4ftgoz6MrF9Y5QusWb5EKe",
	"27vQGMUjXcuraR3NYgNmhN68a5Z8cChyxdY",
	"bjHeUgucPkmKBfFoQ71M8aLwcq48GTvnFx",
	"2GficCngNeGXoBd4EFaE24rS8Qo6GG7QS7D",
	"2NrDzs9z4QZfjtL4BQERMbs3HrPimieVaz4",
	"cCCJ9szQccp7KpCvWTmzMHQX7evAZicH8x",
	"2YXzF6jutzHGckLg2Q5Qe81np4tyRUjQuHq",
	"cumL5gpTygDX9A7NBfRzwMuLgi59dEHuk4",
	"zLH889p2WsMN281mQMYT6Dc2AxsoJxp9t",
	"kGG7bua2SaUm3B9Q2MEEKXpnwozMbBQ5U",
	"A4Jc8ZSVAoQQfhGWRBTazD5L52mBDUKeac",
	"KbGPXEuZRLLUpFUhoMWgJ7q9sEWttJ29ax",
	"YvYCrbz6TJNoUyDybeU8dsQHSQKUKFv4ai",
	"8adTjMnefffco7R9jvKEV5UyfAWufVJHG7",
	"oL8HwXKMfPBJYGLrSA2uMGMsBaMtDSYjhf",
	"7VbqrSuNcFNCsgwdQwHQdhke5EPvyzv96a",
	"kgrpCx6zbqEXsszetUKJt48TU1khsbc8bE",
	"Tt3ftUpQp9DAREvHtPW45K7gCWG6G7SV5q",
	"2fvWAa6JmAXEnWu2ZFVVkUv7F2zfwiyfYiz",
	"yEckRmnMJe4JyAERdtBEgDqEoysTK4XzFW",
	"2gbLZn94HAd61HRE5HT4mpq5Noiag7s3uHY",
	"DM5igZqkjaypXPBNJBAJxb5wjBFx8DrxJb",
	"QzmEkEaUoE4a8TK8xyLJnwkshpvLqTipWX",
	"2Cp5UJ75kD3ibWg3mJJ3o9NUi7pPTRMrZma",
	"2PcQeKm8RhjAMgWVE2PTAMEY3Ey9diFbVse",
	"2XqVkRgorTKCAPZ4Wp7hPHjicvDpk6xSqKJ",
	"x11Wv2bYfW9LmV7TjRs4sUuW1meQpiejSC",
	"MhKtncWDggfFnCmSuaZNcPDc5tejH89sNe",
	"2DjfcKH7TKhpnwyaXjahTAedgJbZt5VtsHL",
	"2bfno4qMcZRQjzAbic69TKK2DNt1rDNbr2z",
	"pobWKFhrtKHtGtFLiYXJD4s6rSGHHrs9Hi",
	"2ZNt8iiDWQvhM9rUjVBfhCbthXqqJdMqQqU",
	"cVUYnASixjP8viCvngHuxRK1iQUd8jzbR1",
	"QKPotkMKHyMam3VcxbmEQ23VLujYCPEktc",
	"2kuaCPDVKeYZBFUvEiY63gLmYEGC81h9u2m",
	"24MHCuxt1muh88rdvUZPesUAMK2J3ELvHvi",
	"C54VisXrMikZRQfdFah2iMYyvTuwvbc5YD",
	"3hbhK6BSLM2JCYAPXEX8yXYA72HyTurcNr",
	"9cgeN9XBhQZ4hnAWW9YmZn7f9aFzDs4EBH",
	"6NjXypCJLHMZWACsYSkaYEAjhzkux871T6",
	"2N59QY3KkSipnxngDWaqwXx4RPjQ9CZWiVt",
	"2bdWDd6GU5ygGp4Cb49uUoFZyK8M5tdUiBm",
	"HTQDDuunbJm2n71odnsaEGzjeZm3K5FgpT",
	"WfgnCGazyERXHG3uK41tv39Pb1BQDRL5jE",
	"qSWMaVVbEvPr1hnVhkDF4qLgin1Js6iNcL",
	"XpXF1QEVB8xdiakXmcL7ddWHRMzn9vfdWi",
	"27VxWNNETMT2zgyNDtdLC1MPDxB5Muapn9w",
	"WEdLzhehQV2hY4sF5HTZz9rezdcyd64sum",
	"2kzWgsWfRsGFBLcZ81vB5MefSUEFfi6KGPj",
	"hwNhnULXK2VNtzURZzAGJomHTuGrj8fiNk",
	"25LdfMWKcf4c1aup27TYyYmTWhRj5AVNRv7",
	"cZtNWh5df2U8iJzZ3jgLqhc4XrmALhpPXc",
	"XkmojuPska7wCvtiTjbYSj9XKxM6ipzN79",
	"2GsRkkoo7gBC1JBXpQU4SdonbjPyJ4B5RyJ",
	"2Bn7qkGf634WTzDvCtTbefrZxMPKM9J8rVK",
	"Sw8txKjkrMmwcKvhaDEkXy8ynwTt85Hr9q",
	"2YG6KYTBXgoDzSzXFMwJT8MyAnqxAK8TPxe",
	"wBJqGdqDdFFW1hkEYjbW78CBoWNte1trLM",
	"f44zDJ7FnV2z8vTUEHcWmsjHiXQWG57j8W",
	"sNZmivMZf4yMod2BMeTD2zVuhSdV4kpWX5",
	"CjKGor7zQdiLzFE8Ki8HNLLv83uoZ5VvuV",
	"obsXPFc5vEYX3U9VvNyjcnFXW5hhwnivKx",
	"Qx59qnE2TJ6HPSF1PDwGTHahUGWp2ZMJZu",
	"VJFTvsnyf5FbEXZVshaQpZ4CpyojaPcf53",
	"Wz5TafrenSzazT3S5NdJ2eZnT3KsbYuNSD",
	"GctVEqovyXMeY3SDfr6Em52D5WF6nJKh8G",
	"KVxMkuSyBWjfavQKxDoqiuQoHiQvJuQD1e",
	"2dHXQFVzPRxNniDZDKoFcVe5HuYVzQz9hB1",
	"wCE4BhH5osH6wf2LXBWwocGp6EWeGodwGg",
	"2Lk1Qq2Lknaa6aYSoHF2gMF2gwTNRHgAudb",
	"28NAqtr9JFzeCJoCtn6TW33DLHSGBDXdBga",
	"2ULzG1HwgyB6XYXpptMn6qD57BWRrHTgDHn",
	"fFMXKFSKdLKCJ847v8JrLGNYyxm26yzBrB",
	"2TAirHnK9Eb7kWfjEmKv3h5FkUPRJ8GG1o1",
	"2YcybfkFNDDD6LS59QGDXWqPonyXTGeKioY",
	"2D4SmhGkch5rx3AXtEMCKuMFvBDt7vrLxei",
	"2cHaD8jmZvwHMTJVnt5EtPJL4qZjBZs9xZF",
	"TZ9DoZJXHxuJPJzkRoxvPFW3qRNivWdy9g",
}
