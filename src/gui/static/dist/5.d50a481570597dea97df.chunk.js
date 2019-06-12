webpackJsonp([5],{TKku:function(e,t){e.exports={common:{"coin-id":"MDL","coin-hours":"Token Hours",usd:"USD",loading:"Loading...",new:"New",load:"Load"},errors:{error:"Error","fetch-version":"Unable to fetch latest release version from Github","incorrect-password":"Incorrect password","error-decrypting":"Error decrypting the wallet","api-disabled":"API disabled","no-wallet":"Wallet does not exist","no-outputs":"No unspent outputs","window-size":"The window is too narrow for the content"},title:{language:"Select Language",wallets:"Wallets",send:"Send",history:"History","buy-coin":"Get MDL",network:"Networking",blockchain:"Blockchain",outputs:"Outputs",transactions:"Transactions","pending-txs":"Pending Transactions",backup:"Backup Wallet",explorer:"MDL Explorer",seed:"Wallet Seed",qrcode:"QR Code",reset:"Reset Password",exchange:"Exchange","select-address":"Select Address","order-history":"Order history"},header:{"syncing-blocks":"Syncing blocks",update1:"Wallet update",update2:"available.",synchronizing:"The wallet is synchronizing. Data you see may not be updated.","pending-txs1":"There are some","pending-txs2":"pending transactions.","pending-txs3":"Data you see may not be updated.",errors:{"no-connections":"No connections active, your client is not connected to any other nodes!","no-backend1":"Cannot reach backend. Please restart the app and/or seek help on our","no-backend2":"Telegram.","no-backend3":"",csrf:"Security vulnerability: CSRF is not working, please exit immediately."}},password:{title:"Enter Password",label:"Password","confirm-label":"Confirm password",button:"Proceed","reset-link":"I forgot my password"},buy:{"deposit-address":"Choose an address to generate a Selected Coin deposit link for:","select-address":"Select address",generate:"Generate","deposit-location":"Deposit Location","deposit-location-desc":"Choose a wallet where you'd like us to deposit your MDL after we receive your Selected Coin.","make-choice":"Make a choice","wallets-desc":"Each time a new wallet and address are selected, a new Selected Coin address is generated. A single MDL address can have up to 5 Selected Coin addresses assigned to it.",send:"Send Selected Coin","send-desc":"Send Selected Coin to the address below. Once received, we will deposit the MDL to a new address in the wallet selected above at the current rate of {{ rate }} MDL/Selected Coin.","fraction-warning":"Only send multiple of the MDL/BTC rate! MDL is sent in whole number; fractional MDL is not sent!",receive:"Receive MDL","receive-desc":"After receiving your Selected Coin, we'll send you your MDL. It may take anywhere between 20 minutes and an hour to receive your MDL.","status-button":"Status:","check-status-button":"Check Status","new-order-button":"New Order"},wizard:{"wallet-desc":'If you don\'t have a MDL wallet, use the generated seed to create a new one. If you already have a wallet, toggle over to "Load Wallet" and enter your seed.',"encrypt-desc":"Increase security of your wallet by encrypting it. By entering a password below, your wallet will be encrypted. Only those with the password will be able access the wallet and remove funds.","hardware-wallet-link":"Using a hardware wallet?","finish-button":"Finish","back-button":"Back",confirm:{title:"Safeguard your seed!",desc:"We want to make sure that you wrote down your seed and stored it in a safe place. If you forget your seed, you WILL NOT be able to recover your MDL wallet!",checkbox:"It\u2019s safe, I swear.",button:"Continue"}},wallet:{"new-address":"New Address","new-addresses":"New Addresses","show-empty":"Show Empty","hide-empty":"Hide Empty",encrypt:"Encrypt Wallet",decrypt:"Decrypt Wallet","decrypt-warning":"Warning: for security reasons, it is not recommended to keep the wallets unencrypted. Caution is advised.",delete:"Delete Wallet",edit:"Edit Wallet",add:"Add Wallet",load:"Load Wallet","encryption-enabled":"Encryption enabled","encryption-disabled":"Encryption disabled","hw-security-warning":'Possible security risk. Access the hardware wallet options (by pressing the "Hardware Wallet" button below the wallets list) while the device is connected for more information.',wallet:"Wallet","hardware-wallet":"Hardware Wallet","delete-confirmation":'WARNING: The wallet "{{ name }}" will be removed from the list. To add it again, you will have to reconnect the device and open the hardware wallet options (by pressing the "Hardware Wallet" button below the wallets list). Do you want to continue?',"delete-confirmation-check":"Yeah, I want to delete the wallet.","max-hardware-wallets-error":"You have already reached the max number of addresses that can be added to the hardware wallet.","add-many-confirmation":"WARNING: If you add too many addresses without using the previous ones or if you use the last ones and not the first ones, some addresses may not be recovered automatically if you try to restore the wallet using the seed (you will have to add them manually). Do you want to continue?",new:{"create-title":"Create Wallet","load-title":"Load Wallet","encrypt-title":"Encrypt Wallet","name-label":"Name","seed-label":"Seed","confirm-seed-label":"Confirm seed","seed-warning":"Remember this seed! Keep it in a safe place. If you forget your seed, you will not be able to recover your wallet!","create-button":"Create","load-button":"Load","cancel-button":"Cancel","12-words":"12 words","24-words":"24 words","generate-12-seed":"Generate 12 word seed","generate-24-seed":"Generate 24 word seed",encrypt:"Encrypt wallet","encrypt-warning":"We suggest that you encrypt each one of your wallets with a password. If you forget your password, you can reset it with your seed. Make sure you have your seed saved somewhere safe before encrypting your wallet.","unconventional-seed-title":"Possible error","unconventional-seed-text":"You introduced an unconventional seed. If you did it for any special reason, you can continue (only recommended for advanced users). However, if your intention is to use a normal system seed, you must delete all the additional text and special characters.","unconventional-seed-check":"Continue with the unconventional seed."},rename:{title:"Rename Wallet","name-label":"Name","cancel-button":"Cancel","rename-button":"Rename"},"add-addresses":{title:"Select quantity","name-quantity":"How many addresses to create","cancel-button":"Cancel","create-button":"Create"},address:{show:"Press to show",copy:"Copy","copy-address":"Copy address",copied:"Copied!",confirm:"Confirm address",outputs:"Unspent outputs",history:"History"}},send:{"synchronizing-warning":"The wallet is still synchronizing the data, so the balance shown may be incorrect. Are you sure you want to continue?","from-label":"Send from","to-label":"Send to","amount-label":"Amount","notes-label":"Notes","wallet-label":"Wallet","addresses-label":"Addresses","invalid-amount":"Please enter a valid amount","addresses-help":"Limit the addresses from where the coins and hours could be sent","all-addresses":"All the addresses of the selected wallet","outputs-label":"Unspent outputs","outputs-help":"Limit the unspent outputs from where the coins and hours could be sent. Only the outputs from the selected addresses are shown","all-outputs":"All the outputs of the selected addresses","available-msg-part1":"With your current selection you can send up to","available-msg-part2":"and","available-msg-part3":"(at least","available-msg-part4":"must be used for the transaction fee).","change-address-label":"Custom change address","change-address-select":"Select","change-address-help":'Address to receive change. If it\'s not provided, it will be chosen automatically. Click on the "Select" link to choose an address from one of your wallets',"destinations-label":"Destinations","destinations-help1":"Destination addresses and their coins","destinations-help2":"Destination addresses, their coins and coin hours","hours-allocation-label":"Automatic coin hours allocation","options-label":"Options","value-label":"Token hours share factor","value-help":"The higher the value, the more token hours will be sent to outputs","preview-button":"Preview","send-button":"Send","back-button":"Back",simple:"Simple",advanced:"Advanced","select-wallet":"Select Wallet"},reset:{"wallet-label":"Wallet","seed-label":"Wallet seed","password-label":"New password (leave empty if you want the wallet not to be encrypted)","confirm-label":"Confirm new password","reset-button":"Reset"},tx:{transaction:"Transaction","confirm-transaction":"Confirm Transaction",from:"From",to:"To",date:"Date",status:"Status",coins:"Tokens",hours:"Hours",id:"Tx ID","show-more":"Show more","hours-moved":"moved","hours-sent":"sent","hours-received":"received","hours-burned":"burned",inputs:"Inputs",outputs:"Outputs",confirmed:"Confirmed",pending:"Pending","current-rate":"Calculated at the current rate"},backup:{"wallet-directory":"Wallet Directory:","seed-warning":"BACKUP YOUR SEED. ON PAPER. IN A SAFE PLACE. As long as you have your seed, you can recover your coins.",desc:"Use the table below to get seeds from your encrypted wallets. <br>To get seeds from unencrypted wallets, open the folder above, open the .wlt files in a text editor and recover the seeds.","close-button":"Close",wallet:"Wallet Label",filename:"Filename",seed:"Seed","show-seed":"Show seed","no-wallets":"No encrypted wallets"},blockchain:{blocks:"Number of blocks",time:"Timestamp of last block",hash:"Hash of last block","current-supply":"Current MDL supply","total-supply":"Total MDL supply","current-coinhour-supply":"Current Token Hours supply","total-coinhour-supply":"Total Token Hours supply"},network:{peer:"Peer",source:"Source","block-height":"Block height","block-height-short":"Block","last-seen":"Last seen","last-received":"Last received","last-sent":"Last sent",in:"Incoming",out:"Outgoing",sources:{default:"Default peer",exchange:"Peer exchange"}},"pending-txs":{timestamp:"Timestamp",txid:"Transaction ID",none:"No pending transactions",my:"Mine",all:"All"},history:{"tx-detail":"Transaction Detail",moving:"Internally moving",moved:"Internally moved",sending:"Sending",sent:"Sent",received:"Received",receiving:"Receiving",pending:"Pending","no-txs":"You have no transaction history","no-txs-filter":"There are no transactions matching the current filter criteria","no-filter":"No filter active (press to select wallets/addresses)",filter:"Active filter: ",filters:"Active filters: ","all-addresses":"All addresses"},teller:{done:"Completed","waiting-confirm":"Waiting for confirmation","waiting-deposit":"Waiting for deposit","waiting-send":"Waiting to send MDL",unknown:"Unknown"},confirmation:{"header-text":"Confirmation","confirm-button":"Yes","cancel-button":"No",close:"Close"},service:{api:{"server-error":"Server error"},wallet:{"not-enough-hours":"Not enough available Coin Hours to perform the transaction!"}},"hardware-wallet":{general:{"default-wallet-name":"New hardware wallet","error-disconnected":"Unable to perform the operation. The hardware wallet is not connected.","simple-error":"Error, Unable to perform the operation.","generic-error":"Unable to perform the operation. Make sure you have connected a valid hardware wallet and that it is not waiting for input.","generic-error-internet":"Unable to perform the operation. Make sure you have connected a valid hardware wallet that is not waiting for input and that you have a connection to the internet/node.","error-incorrect-wallet":"Unable to perform the operation. The connected hardware wallet is different from the expected one.","error-incorrect-pin":"Unable to perform the operation. The PIN you have entered is not correct.",confirm:"Please, confirm the operation in the hardware wallet.","confirm-and-more":"Please, confirm the operation in the hardware wallet and follow the instructions.",close:"Close",cancel:"Cancel",continue:"Continue",completed:"Operation completed successfully.",refused:"The operation failed or was canceled."},errors:{"too-many-inputs":"The transaction has too many inputs and can not be processed. Please contact technical support.","too-many-outputs":"The transaction has too many outputs and can not be processed. Please contact technical support."},"security-warning":{title:"Security warning",text:'The last time this hardware wallet was connected, one or more security warnings were found. We recommend you to open the hardware wallet options (by pressing the "Hardware Wallet" button below the wallets list) while the device is connected and solve the security problems before continuing.',check:"I understand the risks and want to continue",continue:"Continue",cancel:"Cancel"},options:{connecting:"Connecting...",disconnected:"No hardware wallet detected. Please connect a hardware wallet to use this option.","unconfigured-detected-title":"Unconfigured hardware wallet","unconfigured-detected":'A seedless hardware wallet has been detected. Select "Configure automatically" if you want to configure it as a brand new wallet and start using it immediately, or select "Restore backup" if you want to configure it with a previously created seed backup and thus be able to access your funds again.',"configured-detected":"Hardware wallet detected. The device is identified in the wallets list as:","security-warnings-title":"Security warnings","security-warning-title":"Security warning","backup-warning":'You should backup the hardware wallet seed or you could lose access to the funds in case of problems. To do this, select the "Create a backup" option.',"pin-warning":'The connected hardware wallet does not have a PIN. The PIN code protects the hardware wallet in case of loss, theft and hacks. To create a PIN code, select the "Create PIN code" option.',options:"Options:","configure-automatically":"Configure automatically","restore-backup":"Restore backup","create-backup":"Create a backup",wipe:"Wipe the device","confirm-seed":"Confirm seed","create-pin":"Create PIN code","change-pin":"Change PIN code","forgotten-pin1":"If you can not access the wallet because you have forgotten the PIN, you can wipe the hardware wallet and then restore it with the seed by clicking","forgotten-pin2":"here."},"generate-seed":{text:"Before proceeding, you can select the number of words you want the seed to have. The seed is a list of words that can be used to recover access to the coins in case of problems. Both values are safe, so if you do not have a special reason for selecting one or the other, you can leave the default value.",configuring:"Configuring..."},"restore-seed":{text:"Before proceeding, please select the number of words that the seed you want to recover has.","check-text":"You can use this option to enter a seed and check if it is equal to the one in the hardware wallet. Before start, select the number of words the seed you want to check has.",warning:"WARNING: to avoid potential problems, use only seeds created with a hardware wallet from the same brand/model.","error-wrong-word":"Error: The retyped word does not match the one requested by the hardware wallet.","error-invalid-seed":"Error: The seed is not valid. Please be sure to enter the correct words in the correct order.","error-wrong-seed":"Error: The seed is valid but does not match the one in the device."},added:{title:"New Hardware Wallet",configuring:"New hardware wallet detected. Configuring...",done:"Done",added1:"The connected hardware wallet has been added to the wallets list with the following name:",added2:"Now you can check the balance and the addresses of the hardware wallet even when it is not connected. Additionally, you can change the name of the wallet or remove it from the wallets list, if you wish."},wipe:{warning:"WARNING: All the data in the hardware wallet will be deleted. If you do not have a backup, you will not be able to access your funds again.","confirm-delete":"Also remove from the wallets list"},"create-backup":{warning:'WARNING: You can only use this option to backup your hardware wallet seed once. If you decide to continue, you will have to write down a group of words (it is recommended to do it on paper and not on a computer) that will appear on the screen of the hardware wallet and store the list in a safe place. Anyone with access to the word list (the "seed") will be able access the wallet balance, so extreme caution is advised.',instructions:"Write down the word list that appear on the screen of the hardware wallet. Make sure you respect the order and write each word correctly."},"seed-word":{title:"Enter word",info1:"Enter the word indicated in the device",info2:"You will be asked to enter the words of your backup seed in random order, plus a few additional words.",word:"Requested word","error-invalid-word":"The entered word is not valid.","error-loading-words":"Loading the word list. Please wait."},"change-pin":{"pin-mismatch":"Unable to perform the operation. Two PINs you have entered do not match."},"enter-pin":{title:"Enter PIN","title-change-current":"Enter the current PIN","title-change-new":"Enter the new PIN","title-change-confirm":"Confirm the new PIN",instructions:"The PIN layout is displayed on the hardware wallet screen.","instructions-tx":"Enter the PIN to confirm and sign the transaction.","instructions-change":"Please enter a hard-to-guess PIN of between 4 and 8 numbers.",help:"Need help?"},"pin-help":{title:"Help",part1:"When it is necessary to enter the PIN to continue, the hardware wallet screen will display a matrix of 9 boxes with numbers in random order (the order changes each time) and you will be asked to enter the PIN in the software wallet using a matrix of 9 buttons that simply show the symbol #.",part2:'To enter the PIN, look at the position of the PIN numbers in numbers matrix on the screen of the hardware wallet and press the corresponding buttons in the software wallet. For example, if the PIN is "23" and the number 2 is in the upper left, number 3 in the middle of the hardware wallet numbers matrix, press the upper left and middle button in that order in the software wallet.',part3:'If you wish, you can also use the numpad on your keyboard to enter the PIN. However, as in the previous example, if the PIN is "23", you can not simply type "23" with the numpad, but you will have to press the keys that are in the position where the numbers 2 and 3 are shown on the screen of the hardware wallet.'},"create-tx":{title:"Create transaction"},"confirm-address":{title:"Confirm address",instructions:"Please confirm on the hardware wallet if the address is:","short-confirmation":"Address confirmed.",confirmation:'Address confirmed. For security, you can re-show the address in the hardware wallet using the "Confirm address" option, in the menu that you can display by pressing the button at the right of the address balance.'}},"time-from-now":{"few-seconds":"a few seconds ago",minute:"one minute ago",minutes:"{{time}} minutes ago",hour:"one hour ago",hours:"{{time}} hours ago",day:"one day ago",days:"{{time}} days ago"},exchange:{"you-send":"You send","you-get":"You get (approx.)","to-address":"To {{coin}} address",price:"Exchange rate","time-15":"Exchange time","exchange-button":"Exchange","min-amount":"Minimum amount:","max-amount":"Maximum amount:","agree-1":"I agree with","agree-2":"Terms of Use","agree-3":"and","agree-4":"Privacy Policy","powered-by":"Powered by","need-help":"Need help?","support-portal":"Support portal",history:"Order history","order-not-found":"Order not found",status:"Status",exchanging:"Exchanging {{from}} for {{to}}",select:"Select",offline:"Exchange is temporarily offline","problem-connecting":"Unable to connect with the service. Please check your Internet connection and try again later.","invalid-address":"Invalid address.",statuses:{"user-waiting":"Waiting for deposit. Please send {{amount}} {{from}} to the exchange address shown below","user-waiting-info":"The system is waiting for you to make the deposit into the exchange address. The exchange process will start after the deposit is detected and confirmed in the blockchain. If you have already made the deposit, it should be detected shortly.","market-waiting-confirmations":"Waiting for transaction confirmations","market-waiting-confirmations-info":"The deposit has already been detected and the system is waiting for it to be confirmed in the blockchain.","market-confirmed":"Transaction accepted","market-confirmed-info":"The transaction has already been confirmed in the blockchain. Preparing to make the exchange soon.","market-exchanged":"Traded {{from}} for {{to}}","market-exchanged-info":"The exchange has been made. The funds will be transferred to your address in a moment.","market-withdraw-waiting":"Sending {{to}} to your address","market-withdraw-waiting-info":"The process for sending the coins to your address has been initiated.",complete:"Exchange completed!","complete-info":"The funds have been successfully sent to your address.",error:"Error occurred","error-info":"There was an error in the operation, you can find more information below. If you need help, please save all the operation data shown below and contact technical support using the link in the lower right part of this page."},"history-window":{address:"Address",date:"Date"},details:{"exchange-addr":"Exchange address","exchange-addr-tag":"Payment ID or Destination Tag which must be used for the transaction","tx-id":"Transaction ID","order-id":"Order ID","initial-price":"Initial exchange rate","error-msg":"Error message",details:"Details","start-date":"Start date",back:"Back","back-alert":'The operation is still in progress. Do you really want to return to the form? You can see the progress of this operation again by pressing the "Order history" button'}}}}});