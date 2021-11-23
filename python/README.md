# Itzuli API library

* [How to get your api key](https://itzuli.vicomtech.org/api/)
* [All available libraries](https://github.com/Vicomtech/itzuli-api-lib)

# Functions in "Itzuli" package
```
getTranslation(text, fromlang, tolang)
getQuota()
sendFeedback(id, correction, evaluation)
addTag(tagName)
```
An exception will be thrown if any error detected.

# Example

Install the package

> pip install Itzuli

Import:

    from Itzuli import Itzuli

Initialize with your api key:

    itzuliapi = Itzuli("apikey")

Make a translation:

    itzuliapi.getTranslation("kaixo! zer moduz?", "eu", "es")

Get quota:

    itzuliapi.getQuota()

Send feedback:

    itzuliapi.sendFeedback("id from translation, "corrected text", 1)
