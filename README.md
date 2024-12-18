# Opensail

Free and Opensource Regatta Rating System
---

Opensail is a fully transparent revision controlled regatta yacht database using a predictable and deterministic handicap factor.

### Registration


### Development

This project consists of three different components:

- **engine**
- **openfactor**
- **web dashboard**

The **engine** is essentially a replacement for common ci scripts. Due to the complexity of opensail ci tasks, the engine provides a powerful ci tool that avoids the use of multiple shell or python scripts that lack required functionality. 


**openfactor** is a go package containing the code to calculate the opensail openfactor. The package is used by the engine itself, but is abstracted into a separate module.


**web dashboard** is a sveltekit app providing the opensail dashboard. The dashboard is fully independent, all opensail data (ships, teams, etc.) is fetched directly from the latest opensail github release.

Static configuration options are provided via vite environment variables (defaults defined in `.env`).

> [!NOTE]
> Notice that neither project structure nor coding style follows any best practice. The application is specifically designed to fit this use case.