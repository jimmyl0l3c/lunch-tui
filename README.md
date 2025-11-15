# lunch-tui

lunch-tui is a lunch menu scraper with purely text UI printed to the stdout. Expanding the interactivity is planned
in future versions, currently there are only couple CLI arguments you can use to customize the output.

To see all the options, run the utility with `-h` or `--help` argument.

## Migrating from v1

In order to get the same output as with the v1, run the utility with the following arguments:

```bash
lunch-tui dashboard --show-ip
```

## Scrapers

What scrapers should be used can be specified using a JSON config file and `--scrapers` argument. All specified scrapers
will be used and their result will be shown in a single row (that is subject to change).

The JSON configuration should be a list of scraper configs, for example:

```json
[
  { "type": "rozmaryny" },
  {
    "type": "olomouc",
    "title": "Bistro Paulus",
    "menuID": "Bistro-Paulus-6806"
  }
]
```

### Olomouc.cz

Olomouc scraper scrapes the lunch menu section of [olomouc.cz](https://www.olomouc.cz/poledni-menu/).

In order to use it, find a restaurant you are interested in on the site and copy what we call `menuID` (the part of
the URL after `/poledni-menu/`). For example `Bistro-Paulus-6806` in case of `https://www.olomouc.cz/poledni-menu/Bistro-Paulus-6806`.

**Fields:**

- _type:_ mandatory for all scrapers, has to have the value `olomouc`
- _title_: the title displayed in the lunch-tui output
- _menuID_: the ID of the restaurant, see the paragraph above for information how to get it

**Full example:**

```json
{
  "type": "olomouc",
  "title": "Bistro Paulus",
  "menuID": "Bistro-Paulus-6806"
}
```

### Rozmaryny.cz

Rozmaryny scraper scrapes the [Bistro Rozmarýny](https://rozmaryny.cz), it has no additional options, only the mandatory
`type` which has to be `rozmaryny`.

**Full example:**

```json
{ "type": "rozmaryny" }
```

### Creating a new scraper

Since v2, it should be pretty easy to add additional scrapers. Feel free to create PR with new scrapers.

To create a new scraper, create a file in the [./scraping/providers](./scraping/providers) directory with some reasonable
name to describe it. You should create a new struct that has to implement the `scraper.RestaurantScraper` interface.

Use [github.com/gocolly/colly](https://github.com/gocolly/colly) for the scraping. And the `scraper.RetryScrape` instead of
`collector.Visit(url)`, if possible, to ensure the scraping does multiple attempts in case of connectivity issues.

You also need to create a new scraper config in [./scraping/config/providers.go](./scraping/config/providers.go)
(it must embed the `BaseScraperCfg`). Then add a new scraper type to [./scraping/types](./scraping/types/types.go)
and add a new _case_ for it to the unmarshaller of [ScrapersConfig](./scraping/config/config.go).

Last but not least, please document it in this readme in the same manner as the existing scrapers.

**Checklist:**

- Create scraper implementing `scraper.RestaurantScraper` (use `gocolly/colly`).
- Add new scraper type to `ScraperType` enum.
- Create new scraper config that must embed the `BaseScraperCfg`.
- Update the unmarshaller of `ScrapersConfig` to handled the new scraper type.
- Document it in readme
