FROM enakai00/centos:centos6
ENV HOME /root
RUN yum -y install mysql-server
ADD init.sh /usr/local/bin/init.sh
ADD my.cnf /etc/my.cnf
RUN chmod u+x /usr/local/bin/init.sh

EXPOSE 3306
CMD ["/usr/local/bin/init.sh"]
