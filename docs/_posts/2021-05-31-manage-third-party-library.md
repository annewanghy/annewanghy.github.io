---
layout: post
title: "Use Google Tag Manager to manager third-party codes"
date: 2021-06-1 19:47:55 +0800
categories: Analytics Tools
---

such as

1. Google Analytics
2. Google Ads
3. Google Optimize

### How it works

1. first google tagmanager need to install some codes in your side
2. it will generate a variable called `dataLayer` in browser
3. you can use `dataLayer.push()` to put data in this variable
4. use `event` or `pageview` to trigger some tags in google tagmanager
5. tag is been triggered, which means data in dataLayer will been carried by the third-party codes APIs
6. then you can double check API's params and see if it's been send successfully
