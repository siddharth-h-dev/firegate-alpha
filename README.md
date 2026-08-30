# Firegate Alpha
<div align="center">
  <img width="1280" height="640" alt="Firegate Alpha Cover" src="https://github.com/user-attachments/assets/4ce49599-b9e8-4733-85b9-6ee6860812c6" />
</div>
This is an Alpine-based Firewall OS built for Homelabbers that abides by the "Don't Trust Your Firewall" (DTyF) principle.

[Check out DTyF here!](https://github.com/siddharth-h-dev/dont-trust-your-firewall)

WARNING: This is only a firewall, it does not have DHCP or other routing features. Thus, an L3 Managed switch with ip address 172.16.0.2/30 is required for a complete network setup.

## FEATURES:
- **Traffic Shaping**: Built-in Traffic Shaping setup with **CAKE** (Common Applications Kept Enhanced) algorithm having *wash* and other features.
- **IDS/IPS**: Detect (or) Prevent Intrusion with Suricata configued out-of-the-box.
- **Tor routing**: Certain VLAN traffic can be routed through Tor for anonymity on the internet.
- **Unbound with Ad-blocking**: Unbound DNS with a blocklist for ads, trackers, malware and nsfw.
- **Easy-to-use Web GUI**: Web GUI designed for easy configuration without the need to edit configuration files manually.
- **GitOps config management**: Configuration is powered by GitOps.
- **Abides by DTyF**: Firegate Alpha is built such that you can implement the DTyF principle in your network easily.

## ADVANTAGES:
- **Standalone Firewall**: As a separate L3 Switch will handle DHCP, Inter-VLAN Routing, etcetra, Firegate will be efficient at its core purpouse.
- **Homelabber-Centric**: Built with the homelabbers in mind.
- **Minimal Setup**: Core features are already setup for you.
- **Secure**: Definetly more secure than a standard home router.
- **Super Lightweight**: Alpine Linux is the base for this OS.
- **Easy rollbacks**: You can easily revert the config if it is broken.
- **Prepared for a compromise**: If you follow the DTyF principle, you network will be prepared for a firewall compromise.

Core features developed. Web GUI still in development.

