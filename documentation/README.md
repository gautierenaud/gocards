# Description

I'll try to list all the things I want to do in this file.

# Tasks
## Doings

Things that I'm likely to tackle first.

* send a notification when the db is updated

## TODOs

* card modifiers (foils, special treatment, promo, set, art, ...)
    * it could be a combination of modifiers
-> how to compare cards (I want the cards with the same set of modifiers to be grouped together)
    must be compatible with sql queries
    -> first step: set trigram and card number are enough to differentiate the cards
        -> no support for foil, promo or anything else

* sections
    * all collections section
    * section by deck/tags? ("regroupments")

* import/export
    * through interface (drag and drop file)
    * display cards from a set and click on the one we have
    * preview on import (so we can fix the set/art of the card before committing)

* upload custom images (signed cards, proxy, ...)
    * change image of existing card

* background
    * like the cards are on a desk
    * like the cards are in a binder

* support flip cards

* local first app -> download the images for the future

* zoom level

* fix image retrieval -> "Reprieve" has image of "Graceful Reprieve"

* order
    * by set + card number in the set, name, ...

* add custom tag on each card (and list by tags)
    * cards in decks, cards in specific decks, ...
    * where the card is physically stored

* deck's history

* what card do I have?
    * get a list of card (from scryfall) and compare them with local db
    * passthrough request to scryfall?

## Done

* all collection.
* import cards (through cli first)
* display: round the edges of the cards
* image with specific language
* card counts (probably figure out the modifiers first)