# Ansible Playbook: Attack Box

[![Molecule Test](https://github.com/l50/ansible-attack-box/actions/workflows/molecule.yaml/badge.svg)](https://github.com/l50/ansible-attack-box/actions/workflows/molecule.yaml)
[![Pre-Commit](https://github.com/l50/ansible-attack-box/actions/workflows/pre-commit.yaml/badge.svg)](https://github.com/l50/ansible-attack-box/actions/workflows/pre-commit.yaml)
[![License](https://img.shields.io/github/license/l50/ansible-attack-box?label=License&style=flat&color=blue&logo=github)](https://github.com/l50/ansible-attack-box/blob/main/LICENSE)

This playbook provisions a system with various tools that are
useful for offensive security work.

---

## Setup and execution

1. Install dependencies:

   ```bash
   sudo apt install -y python3-pip libssl-dev
   python3 -m pip install molecule ansible molecule[docker,lint]
   ```

1. Install galaxy roles and collections

   ```bash
   ansible-galaxy install -r requirements.yml
   ```

1. Run the playbook

   ```bash
   ansible-playbook \
       --connection=local \
       --inventory 127.0.0.1, \
       --limit 127.0.0.1 attack-box.yaml
   ```

---

## Testing

To test changes made to this role, run the following commands:

```bash
# If on an Apple Silicon machine:
if [[ "$(uname -a | awk '{ print $NF }')" == "arm64" ]]; then
  export DOCKER_DEFAULT_PLATFORM=linux/arm64
fi
molecule create
molecule converge
molecule idempotence
# If everything passed, tear down the docker container spawned by molecule:
molecule destroy
```
