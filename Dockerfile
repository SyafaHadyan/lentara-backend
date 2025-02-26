FROM archlinux:latest
# CMD ["sudo", "pacman", "-Syu"]
# CMD ["pacman-key", "--init"]
# CMD ["sudo", "pacman-key", "--populate", "archlinux"]
# CMD ["sudo", "pacman", "-S", "archlinux-keyring"]
# CMD ["sudo", "pacman", "-S", "go", "mariadb", "mariadb-clients", "mariadb-server", "git", "kitty-terminfo"]
# CMD ["sudo", "mariadb-install-db", "--user=mysql", "--basedir=/usr", "--datadir=/var/lib/mysql"] #From Arch Wiki before starting mariadb (with systemd)
# CMD ["sudo", "systemctl", "start", "mariadb"]
# CMD ["sudo", "systemctl", "enable", "mariadb"]
# CMD ["git", "clone", "https://github.com/SyafaHadyan/lentara-backend.git"]
# CMD ["touch", "lentara-backend/.env"]

RUN pacman -Syu --noconfirm && \
    pacman-key --init && \
    pacman-key --populate archlinux && \
    pacman -S --noconfirm archlinux-keyring && \
    pacman -S --noconfirm go mariadb mariadb-clients mariadb-server git kitty-terminfo && \
    mariadb-install-db --user=mysql --basedir=/usr --datadir=/var/lib/mysql && \
    git clone https://github.com/SyafaHadyan/lentara-backend.git && \
    touch lentara-backend/.env

CMD sh -c "systemctl start mariadb && systemctl enable mariadb"

EXPOSE 8080
EXPOSE 3306

