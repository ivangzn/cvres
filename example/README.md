# Profile data file

The resume profile data file format will ultimately depend if it's either a
[JSON](https://en.wikipedia.org/wiki/JSON) or
[YAML](https://en.wikipedia.org/wiki/YAML) file.

The structure of the content is as follows:

## Person

- **name**: The full name of the individual.
- **role**: The professional title or role of the individual.

## Contact

- **email**: The email address of the individual.
- **location**: The geographical location of the individual.
- **linkedin**: The LinkedIn profile URL of the individual.

## Sections

A list of sections that make up the resume.
Each section contains the following fields:

- **title**: The title of the section (e.g., "Profile", "Professional Experience").
- **articles**: A list of articles or entries within the section. Each article can have the following fields:
  - **desc**: A description or summary of the article.
  - **what**: The title or main subject of the article (e.g., job title, course name).
  - **where**: The organization or location associated with the article (e.g., company name, university).
  - **when**: The time period associated with the article (e.g., start and end dates).
  - **list**: A list of items providing additional details.
  - **full-list**: A list of short items, usually printed in columns.
