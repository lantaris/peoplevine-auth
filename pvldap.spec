%define debug_package %{nil}
%undefine _missing_build_ids_terminate_build

Name:    pvldap
Version: 0.0.1
Release: 1%{?dist}
Summary: PeopleVine auth proxy server
License: Private
Source0:  pvldap.tgz
BuildRequires: golang
BuildRequires: git
BuildRequires: systemd

%description
 PeopleVine ldap auth proxy server

%prep
%setup -q -n %{name}

%build

export GOPATH=$(pwd) &&\
cd %{name} &&\
go build .

%install
install -d %{buildroot}%{_sbindir}
install -d %{buildroot}%{_unitdir}
install -d %{buildroot}%{_sysconfdir}
install -p -m 0755 %{name}/%{name} %{buildroot}%{_sbindir}
install -p -m 0644 systemd/%{name}.service %{buildroot}%{_unitdir}/%{name}.service

%post
%systemd_post %{name}.service

%preun
%systemd_preun %{name}.service

%postun
%systemd_postun_with_restart %{name}.service

%files
%doc README.md
%{_sbindir}/*
%{_unitdir}/*

%changelog
* Mon Jan 23 2023 Eduard Ponomarenko <gotik@mail.com> - 0.0.1-1
- initial version

