# Opensail

Free and Opensource Regatta Rating System

![Opensail Icon](/static/favicon.png)

Opensail is a transparent, revision-controlled regatta yacht database offering a predictable and deterministic factor for handicap regattas.


All data is stored in this repository and managed through the github ecosystem. Hosting and deployment are provided by the [battleshiper](https://battleshiper.dev) platform.


For ship owners that want to register or updater their sailing crew or regatta vessel, see [registration](/REGISTER.md).


If you have suggestions, feature requests, or encounter any issues, feel free to open a github issue or contact us at [contact@osail.ch](mailto:contact@osail.ch).



### Openfactor

Openfactor is the algorithm used to calculate the handicap factor for each vessel, based on ship specifications provided during registration. The factor is calibrated using empirical constants, which may be adjusted over time based on feedback and insights from the collected data.


The full implementation of the algorithm is available in the `openfactor/openfactor.go` file within this repository.


### Versioning

Opensail uses git tags for version control; all versions strictly follow the semver format.


The core opensail version is automatically incremented at the patch level every time the `push-register` workflow generates a new version.
This workflow automatically creates a tag and commits the generated api data into `static/api/`. As a result, for every new version an additional commit with the raw api data is executed.

As a static reference, each published version is also packed into a github release asset.


The openfactor module uses a separate versioning system, also managed by tags prefixed with `openfactor/`. This tag must be manually updated when openfactor is modified.


### Development

This project consists of three different components:

- **engine**
- **openfactor**
- **web dashboard**

The **engine** is essentially a replacement for common ci scripts. Due to the complexity of opensail ci tasks, the engine provides a powerful ci tool that avoids the use of multiple shell or python scripts that lack required functionality. The engine is operated by github workflows which validate and process data in `register/**`.


**openfactor** is a go package containing the code to calculate the opensail openfactor. The package is used by the engine itself, but is abstracted into a separate module.


**web dashboard** is a sveltekit app providing the opensail dashboard. All raw data (ships, teams, etc.) is inserted into the `static/api/` by the ci engine, this means the data is treated as static assets of the web app and therefore served via the underlying battleshiper cdn.

> [!NOTE]
> Notice that neither project structure nor coding style follows any best practice. The application is specifically designed to fit this use case.