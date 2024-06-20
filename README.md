# usher-ifcb

File organizer for IFCB data files. usher-ifcb monitors a directory
for newly created files (usually transferred by IFCB computers) and
organizes them into the standard `YYYY/DYYYYMMDD` based on dates
parsed from the filename (expected format `DYYYYMMDDTHHMMSS_IFCBNNN.ext`).
A root directory containing multiple subdirectories which map to
distinct target directories can also be montiored by a single
usher-ifcb process by using usher's `rootpathmappings` configuration.

usher-ifcb uses the [usher](https://github.com/axiom-data-science/usher)
project, see that project's README for more information.
