
numid query gov proposals
numid tx gov submit-legacy-proposal param-change ./proposal.json --from alice
numid tx gov vote 1 yes --from alice
numid query gov proposals
numid query params subspace numi identityVerifiers

