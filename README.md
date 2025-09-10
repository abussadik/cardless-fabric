# Cardless Domestic Payments on Hyperledger Fabric

This repository contains datasets, evaluation assets, and code skeletons for a permissioned, cardless payment prototype on Hyperledger Fabric.

## Structure
- datasets/: synthetic data (1,000 users; 50 merchants; 10,000 tx)
- evaluation/: performance curves + charts (synthetic)
- usability/: SUS survey template
- chaincode/payments-go/: Go chaincode skeleton
- app/server/: Node.js Fabric Gateway backend skeleton
- network/: notes for using fabric-samples test-network
- docker/: compose for app layer

## Quickstart
See `network/README.md` then run the backend in `app/server`.

## Upload to GitHub
```
git init
git add .
git commit -m "Initial commit"
git branch -M main
git remote add origin git@github.com:<YOUR-USERNAME>/cardless-fabric.git
git push -u origin main
```
