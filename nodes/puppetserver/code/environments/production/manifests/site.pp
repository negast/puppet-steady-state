## site.pp ##
node default {

}

node 'ubuntu1' {
  include chaosblade
  include apache

#   chaosexperiment_cpu { 'cpuload1':
#       ensure   => 'present',
#       load     => 99,
#       climb    => 60,
#       timeout  => 600,
#    }


#   chaosexperiment_process { 'stopp1':
#     ensure           => 'present',
#     type             => 'process_stop',
#     process_cmd      => 'apache2',
#     timeout          => 60,
#   }

}
