<p align="center">
  <img src="https://i.imgur.com/ClLeIeK.gif" alt="Zero Two" />
</p>
## About

This was a deprecated project forked and updated by [xqyet](https://github.com/xqyet). I added a number of stability, logging, and 
performance improvements for my personal use.
- **Webhook Logging**
  - log request attempts, success, failures, and rate-limit hits to Discord or Telegram.
  - variable webhooks for debug logging and successful claim notifications - includes bearer tail display.
  - includes status code, request time, bearer ID, and some additional error logging.
- **Dynamic Backoff & Claim Requests**
  - delay scaling based on recent `429 Too Many Requests`.
  - delay customization per request cycle.
  - conditional added to avoid synchronized rate-limiting across accounts.
  - 
## How to Update
The original MCsniperGO repo contains a deprecated OAuth client ID used for the Microsoft device login flow. 
You can find it in the mc/account.go file in this line: 
`client_id := "00000000402b5328"`
This is blacklisted by Microsoft, so if you want to automatically pull bearer tokens and track
your own OAuth sessions, do the following: 

### Register a New Microsoft Azure App

1. Visit: [https://portal.azure.com](https://portal.azure.com)
2. Navigate to **Azure Active Directory > App registrations**
3. Click **"New registration"** and fill out:
  - **Name**: Anything (e.g., `MCsniperGO`)
  - **Supported account types**:  
  "Accounts in any organizational directory and personal Microsoft accounts"
  - **Redirect URI** (optional):  
    `https://login.microsoftonline.com/common/oauth2/nativeclient`

4. Click **Register** and copy your **Application (client) ID**

### Replace It in Code

In `mc/account.go`, replace:

```go
client_id := "00000000402b5328"
 ```
with

```go
client_id := "your-new-client-id"
 ```



## How to Use

- [Install go](https://go.dev/dl/)
- Download or clone MCsniperGO repository 
- open MCsniperGO folder in your terminal / cmd
- put your prename accounts (no claimed username) in [`gc.txt`](#accounts-formatting) and your normal accounts in [`ms.txt`](#accounts-formatting)
- put proxies into `proxies.txt` in the format `user:pass@ip:port` (there should NOT be 4 `:` in it as many proxy providers provide it as)
- run `go run ./cmd/cli`
- enter username + [claim range](#claim-range)
- wait, and hope you claim the username!

## Claim Range
Use the following Javascript bookmarklet in your browser to obtain the droptime while on `namemc.com/search?q=<username>`:

```js
javascript:(function(){function parseIsoDatetime(dtstr) {
    return new Date(dtstr);
};

startElement = document.getElementById('availability-time');
endElement = document.getElementById('availability-time2');

start = parseIsoDatetime(startElement.getAttribute('datetime'));
end = parseIsoDatetime(endElement.getAttribute('datetime'));

para = document.createElement("p");
para.innerText = Math.floor(start.getTime() / 1000) + '-' + Math.ceil(end.getTime() / 1000);

endElement.parentElement.appendChild(para);})();

```

If 3name.xyz has a lower length claim range for a username I would recommend using that, you can get the unix droptime range with this bookmarklet on `3name.xyz/name/<name>`

```js
javascript: (function() {
    startElement = document.getElementById('lower-bound-update');
    endElement = document.getElementById('upper-bound-update');
  
  	if (startElement === null) {
    	startElement = 0;
    } else {
      startElement = startElement.getAttribute('data-lower-bound')
    }
  
  
    para = document.createElement("p");
    para.innerText = Math.floor(Number(startElement) / 1000) + '-' + Math.ceil(Number(endElement.getAttribute('data-upper-bound')) / 1000);
    endElement.parentElement.appendChild(para)
})()
```

## accounts formatting

- place in `gc.txt` or `ms.txt` depending on their account type.
  - `gc.txt` is for accounts without usernames
  - `ms.txt` is for accounts that already have usernames on them
- a bearer token can be obtained by following  [this guide](https://kqzz.github.io/mc-bearer-token/)

```txt
EMAIL:PASSWORD
BEARER
```

### Example accounts file

```txt
kqzz@gmail.com:SecurePassword3
teun@example.com:SafePassword!
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```
> This will load 3 accounts into the sniper, two of which are supplied with email / password. The last is loaded by bearer token, and will last 24 hours (the sniper will show the remaining time).

> Their account types are determined by if they are placed in `gc.txt` or `ms.txt`.

## understanding logs

Each request made to change your username will return a 3 digit HTTP status code, the meanings are as follows:

- 400 / 403: Failed to claim username (will continue trying)
- 401: Unauthorized (restart claimer if it appears)
- 429: Too many requests (add more proxies if this occurs frequently)
