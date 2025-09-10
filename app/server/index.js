const express = require('express');
const { Gateway, Wallets } = require('fabric-network');
const fs = require('fs'); const path = require('path');
const app = express(); app.use(express.json());
const CCP = path.join(__dirname, 'connection.json'); const CHANNEL='domestic-payments'; const CCNAME='payments';
async function submit(id,fcn,args){ const ccp=JSON.parse(fs.readFileSync(CCP,'utf8'));
  const wallet=await Wallets.newFileSystemWallet(path.join(__dirname,'wallet'));
  const gateway=new Gateway(); await gateway.connect(ccp,{wallet,identity:id,discovery:{enabled:true,asLocalhost:true}});
  const network=await gateway.getNetwork(CHANNEL); const contract=network.getContract(CCNAME);
  const res=await contract.submitTransaction(fcn,...args); await gateway.disconnect(); return res.toString(); }
app.post('/wallets', async (req,res)=>{ const {ownerID,walletID,as}=req.body;
  try{ await submit(as,'CreateWallet',[ownerID,walletID]); res.json({ok:true}); }catch(e){ res.status(400).json({error:e.message}); }});
app.post('/payments', async (req,res)=>{ const {txID,type,fromWallet,toWallet,currency,amount,as}=req.body;
  try{ await submit(as,'ProcessPayment',[txID,type,fromWallet,toWallet,currency,String(amount)]); res.json({ok:true,txID}); }
  catch(e){ res.status(400).json({error:e.message}); }});
app.listen(8080, ()=> console.log('API on :8080'));
