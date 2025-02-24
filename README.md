# Privateer Plugin for GitHub Repos

This wireframe is designed to quickly get your service pack repository up to speed!

Privateer's plugin architecture relies on some key elements being present at the top level
of the service pack, all of which are provided along with example code in this repo.

Simply fork or clone this repo and start adjusting the tests to build your own service pack!

Based on the Open Source Project Security (OSPS) Baseline, here is a consolidated checklist of all rules across various categories, including descriptions of what each baseline implies to check.

## Level 1 Criteria Status

| ID         | REST | GraphQL | Criteria |
| ---------- | ---- | ------- | ----------- |
| OSPS-AC-01 |      | x       | The project's version control system MUST require multi-factor authentication for collaborators modifying the project repository settings or accessing sensitive data. |
| OSPS-AC-02 | x    |         | The project's version control system MUST restrict collaborator permissions to the lowest available privileges by default. |
| OSPS-AC-03 |      | x       | The project's version control system MUST prevent unintentional direct commits against the primary branch. |
| OSPS-AC-04 |      | x       | The project's version control system MUST prevent unintentional deletion of the primary branch. |
| OSPS-BR-01 |      |         | The project's build and release pipelines MUST NOT permit untrusted input that allows access to privileged resources. |
| OSPS-BR-03 | x    |         | Any websites and version control systems involved in the project development MUST be delivered using SSH, HTTPS, or other encrypted channels. |
| OSPS-BR-09 |      |         | Any websites or other services involved in the distribution of released software assets MUST be delivered using HTTPS or other encrypted channels. |
| OSPS-DO-03 | x    |         | The project documentation MUST provide user guides for all basic functionality. |
| OSPS-DO-05 |      | x       | The project documentation MUST include a mechanism for reporting defects. |
| OSPS-GV-02 |      | x       | The project MUST have one or more mechanisms for public discussions about proposed changes and usage obstacles. |
| OSPS-GV-03 |      | x       | The project documentation MUST include an explanation of the contribution process. |
| OSPS-LE-02 |      | x       | The license for the source code MUST meet the OSI Open Source Definition or the FSF Free Software Definition. |
| OSPS-LE-03 |      | x       | The license for the source code MUST be maintained in a standard location within the project's repository. |
| OSPS-LE-04 | x    |         | The license for the released software assets MUST meet the OSI Open Source Definition or the FSF Free Software Definition. |
| OSPS-QA-01 | x    |         | The project's source code MUST be publicly readable and have a static URL. |
| OSPS-QA-02 |      |         | The version control system MUST contain a publicly readable record of all changes made, who made the changes, and when the changes were made. |
| OSPS-VM-05 |      |         | The project publishes contacts and process for reporting vulnerabilities. |

## Level 2 Criteria Status

| ID         | REST | GraphQL | Criteria |
| ---------- | ---- | ------- | ----------- |
| OSPS-AC-05 | x    |         | The project's permissions in CI/CD pipelines MUST be configured to the lowest available privileges except when explicitly elevated. |
| OSPS-BR-02 |      |         | All releases and released software assets MUST be assigned a unique version identifier for each release intended to be used by users. |
| OSPS-BR-04 |      |         | All released software assets MUST be created with consistent, automated build and release pipelines. |
| OSPS-BR-05 |      |         | All build and release pipelines MUST use standardized tooling where available to ingest dependencies at build time. |
| OSPS-BR-06 |      |         | All releases MUST provide a descriptive log of functional and security modifications. |
| OSPS-BR-08 |      |         | All released software assets MUST be signed or accounted for in a signed manifest including each asset's cryptographic hashes. |
| OSPS-BR-10 |      |         | Any websites, API responses, or other services involved in release pipelines MUST be fetched using SSH, HTTPS, or other encrypted channels. |
| OSPS-DO-12 |      |         | The project documentation MUST contain instructions to verify the integrity and authenticity of the release assets, including the expected identity of the person or process authoring the software release.      |
| OSPS-DO-13 |      |         | The project documentation MUST include a descriptive statement about the project's intended scope and usage. |
| OSPS-DO-14 |      |         | The project documentation MUST include a list of all project dependencies, including their licenses and versions. |
| OSPS-DO-15 |      |         | The project documentation MUST include a list of all project maintainers and their roles. |
| OSPS-GV-04 |      |         | The project MUST have a documented governance model that defines the decision-making process and roles within the project. |
| OSPS-GV-05 |      |         | The project MUST have a code of conduct that defines the expected behavior of contributors and the process for reporting and handling unacceptable behavior. |
| OSPS-LE-05 |      |         | The project MUST have a documented process for handling license changes, including obtaining contributor agreement for the changes. |
| OSPS-LE-06 |      |         | The project MUST have a documented process for handling contributor license agreements (CLAs) or developer certificates of origin (DCOs), if applicable. |
| OSPS-QA-03 |      |         | The project MUST have a continuous integration system that automatically builds and tests the project's code upon changes. |
| OSPS-QA-04 |      |         | The project MUST have a policy for code reviews that includes the requirement for at least one independent reviewer to approve changes before they are merged. |
| OSPS-QA-05 |      |         | The project MUST have a policy for handling reported defects that includes a process for prioritizing and addressing them in a timely manner. |
| OSPS-SA-01 |      |         | The project MUST perform regular security assessments, including code reviews and vulnerability scanning, to identify and address potential security issues. |
| OSPS-SA-02 |      |         | The project MUST have a documented process for managing security vulnerabilities, including a mechanism for receiving and disclosing vulnerabilities, and a process for issuing security advisories and patches. |
| OSPS-VM-06 |      |         | The project MUST have a documented process for managing dependencies, including a mechanism for tracking and updating dependencies, and a process for handling vulnerabilities in dependencies. |
| OSPS-VM-07 |      |         | The project MUST have a policy for handling end-of-life (EOL) for project releases, including a process for communicating EOL to users and a process for handling security vulnerabilities in EOL releases. |

## Level 3 Criteria Status

| ID           | REST | GraphQL | Criteria |
| ------------ | ---- | ------- | ---------- |
| OSPS-AC-06   |      |         | The project MUST have a documented process for onboarding and offboarding collaborators, including access management and permissions review. |
| OSPS-BR-07   |      |         | The project MUST have a documented process for handling security vulnerabilities in build and release pipelines, including a mechanism for reporting and addressing vulnerabilities. |
| OSPS-DO-16   |      |         | The project documentation MUST include a security policy that defines the project's approach to security, including roles and responsibilities, and a process for reporting and handling security issues. |
| OSPS-GV-06   |      |         | The project MUST have a documented process for managing project assets, including source code, documentation, and other resources, with defined roles and responsibilities for asset management. |
| OSPS-LE-07   |      |         | The project MUST have a documented process for handling trademark and branding issues, including guidelines for the use of the project's name, logo, and other branding elements. |
| OSPS-QA-06   |      |         | The project MUST have a documented process for managing quality assurance, including testing strategies, code review practices, and release criteria, to ensure the quality and stability of the project. |
| OSPS-SA-03   |      |         | The project MUST have a documented process for conducting threat modeling and risk assessments to identify and mitigate potential security risks in the project's design and implementation. |
| OSPS-VM-08   |      |         | The project MUST have a documented process for managing security incidents, including a mechanism for reporting incidents, a process for incident response and recovery, and a process for communicating with stakeholders. |
