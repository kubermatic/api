. as $doc |

$crdfiles[0] as $crd |
$crd.spec.group as $group |
$crd.spec.versions[] | select(.name==$version) | .schema.openAPIV3Schema as $schema |

$doc | .paths += {
   ("/" + $group + "/" + $filename): {
      POST: {
         requestBody: {
            required: true,
            content: {
               "application/json": {
                  schema: $schema
               }
            }
         }
      }
   }
}
