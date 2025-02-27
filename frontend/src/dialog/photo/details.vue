<template>
  <div class="p-tab p-tab-photo-details">
    <v-container fluid>
      <v-form ref="form" lazy-validation
              dense class="p-form-photo-details-meta" accept-charset="UTF-8"
              @submit.prevent="save">
        <v-layout row wrap align-top fill-height>
          <v-flex
              class="p-photo pa-2"
              xs12 sm4 md2
          >
            <v-card tile
                    class="ma-1 elevation-0"
                    :title="model.Title">
              <v-img v-touch="{left, right}"
                     :src="model.thumbnailUrl('tile_500')"
                     aspect-ratio="1"
                     class="accent lighten-2 elevation-0 clickable"
                     @click.exact="openPhoto()"
              >
              </v-img>

            </v-card>
          </v-flex>
          <v-flex xs12 sm8 md10 fill-height>
            <v-layout row wrap>
              <v-flex xs12 class="pa-2">
                <v-text-field
                    v-model="model.Title"
                    :append-icon="model.TitleSrc === 'manual' ? 'check' : ''"
                    :disabled="disabled"
                    :rules="[textRule]"
                    hide-details
                    :label="$gettext('Title')"
                    placeholder=""
                    color="secondary-dark"
                    browser-autocomplete="off"
                    class="input-title"
                ></v-text-field>
              </v-flex>

              <v-flex xs4 md1 pa-2>
                <v-autocomplete
                    v-model="model.Day"
                    :append-icon="model.TakenSrc === 'manual' ? 'check' : ''"
                    :disabled="disabled"
                    :error="invalidDate"
                    :label="$gettext('Day')"
                    browser-autocomplete="off"
                    hide-details hide-no-data
                    color="secondary-dark"
                    :items="options.Days()"
                    class="input-day"
                    @change="updateTime">
                </v-autocomplete>
              </v-flex>
              <v-flex xs4 md1 pa-2>
                <v-autocomplete
                    v-model="model.Month"
                    :append-icon="model.TakenSrc === 'manual' ? 'check' : ''"
                    :disabled="disabled"
                    :error="invalidDate"
                    :label="$gettext('Month')"
                    browser-autocomplete="off"
                    hide-details hide-no-data
                    color="secondary-dark"
                    :items="options.MonthsShort()"
                    class="input-month"
                    @change="updateTime">
                </v-autocomplete>
              </v-flex>
              <v-flex xs4 md2 pa-2>
                <v-autocomplete
                    v-model="model.Year"
                    :append-icon="model.TakenSrc === 'manual' ? 'check' : ''"
                    :disabled="disabled"
                    :error="invalidDate"
                    :label="$gettext('Year')"
                    browser-autocomplete="off"
                    hide-details hide-no-data
                    color="secondary-dark"
                    :items="options.Years()"
                    class="input-year"
                    @change="updateTime">
                </v-autocomplete>
              </v-flex>

              <v-flex xs6 md2 class="pa-2">
                <v-text-field
                    v-model="time"
                    :append-icon="model.TakenSrc === 'manual' ? 'check' : ''"
                    :disabled="disabled"
                    :label="model.timeIsUTC() ? $gettext('Time UTC') : $gettext('Local Time')"
                    browser-autocomplete="off"
                    hide-details
                    return-masked-value mask="##:##:##"
                    color="secondary-dark"
                    class="input-local-time"
                ></v-text-field>
              </v-flex>

              <v-flex xs6 sm6 md6 class="pa-2">
                <v-autocomplete
                    v-model="model.TimeZone"
                    :disabled="disabled"
                    :label="$gettext('Time Zone')"
                    browser-autocomplete="off"
                    hide-details hide-no-data
                    color="secondary-dark"
                    item-value="ID"
                    item-text="Name"
                    :items="options.TimeZones()"
                    class="input-timezone"
                    @change="updateTime">
                </v-autocomplete>
              </v-flex>

              <v-flex xs12 sm8 md4 class="pa-2">
                <v-autocomplete
                    v-model="model.Country"
                    :append-icon="model.PlaceSrc === 'manual' ? 'check' : ''"
                    :disabled="disabled"
                    :readonly="!!(model.Lat || model.Lng)"
                    :label="$gettext('Country')" hide-details
                    hide-no-data
                    browser-autocomplete="off"
                    color="secondary-dark"
                    item-value="Code"
                    item-text="Name"
                    :items="countries"
                    class="input-country">
                </v-autocomplete>
              </v-flex>

              <v-flex xs4 md2 class="pa-2">
                <v-text-field
                    v-model="model.Altitude"
                    :disabled="disabled"
                    hide-details
                    browser-autocomplete="off"
                    :label="$gettext('Altitude (m)')"
                    placeholder=""
                    color="secondary-dark"
                    class="input-altitude"
                ></v-text-field>
              </v-flex>

              <v-flex xs4 sm6 md3 class="pa-2">
                <v-text-field
                    v-model="model.Lat"
                    :append-icon="model.PlaceSrc === 'manual' ? 'check' : ''"
                    :disabled="disabled"
                    hide-details
                    browser-autocomplete="off"
                    :label="$gettext('Latitude')"
                    placeholder=""
                    color="secondary-dark"
                    class="input-latitude"
                ></v-text-field>
              </v-flex>

              <v-flex xs4 sm6 md3 class="pa-2">
                <v-text-field
                    v-model="model.Lng"
                    :append-icon="model.PlaceSrc === 'manual' ? 'check' : ''"
                    :disabled="disabled"
                    hide-details
                    browser-autocomplete="off"
                    :label="$gettext('Longitude')"
                    placeholder=""
                    color="secondary-dark"
                    class="input-longitude"
                ></v-text-field>
              </v-flex>

              <v-flex xs12 md6 pa-2 class="p-camera-select">
                <v-select
                    v-model="model.CameraID"
                    :append-icon="model.CameraSrc === 'manual' ? 'check' : ''"
                    :disabled="disabled"
                    :label="$gettext('Camera')"
                    browser-autocomplete="off"
                    hide-details
                    color="secondary-dark"
                    item-value="ID"
                    item-text="Name"
                    :items="cameraOptions"
                    class="input-camera">
                </v-select>
              </v-flex>

              <v-flex xs6 md3 class="pa-2">
                <v-text-field
                    v-model="model.Iso"
                    :disabled="disabled"
                    hide-details
                    browser-autocomplete="off"
                    label="ISO"
                    placeholder=""
                    color="secondary-dark"
                    class="input-iso"
                ></v-text-field>
              </v-flex>

              <v-flex xs6 md3 class="pa-2">
                <v-text-field
                    v-model="model.Exposure"
                    :disabled="disabled"
                    hide-details
                    browser-autocomplete="off"
                    :label="$gettext('Exposure')"
                    placeholder=""
                    color="secondary-dark"
                    class="input-exposure"
                ></v-text-field>
              </v-flex>

              <v-flex xs12 md6 pa-2 class="p-lens-select">
                <v-select
                    v-model="model.LensID"
                    :append-icon="model.CameraSrc === 'manual' ? 'check' : ''"
                    :disabled="disabled"
                    :label="$gettext('Lens')"
                    browser-autocomplete="off"
                    hide-details
                    color="secondary-dark"
                    item-value="ID"
                    item-text="Name"
                    :items="lensOptions"
                    class="input-lens">
                </v-select>
              </v-flex>

              <v-flex xs6 md3 class="pa-2">
                <v-text-field
                    v-model="model.FNumber"
                    :disabled="disabled"
                    hide-details
                    browser-autocomplete="off"
                    :label="$gettext('F Number')"
                    placeholder=""
                    color="secondary-dark"
                    class="input-fnumber"
                ></v-text-field>
              </v-flex>

              <v-flex xs6 md3 class="pa-2">
                <v-text-field
                    v-model="model.FocalLength"
                    :disabled="disabled"
                    hide-details
                    browser-autocomplete="off"
                    :label="$gettext('Focal Length')"
                    placeholder=""
                    color="secondary-dark"
                    class="input-focal-length"
                ></v-text-field>
              </v-flex>

              <v-flex xs12 sm6 md3 class="pa-2">
                <v-textarea
                    v-model="model.Details.Subject"
                    :append-icon="model.Details.SubjectSrc === 'manual' ? 'check' : ''"
                    :disabled="disabled"
                    :rules="[textRule]"
                    hide-details
                    browser-autocomplete="off"
                    auto-grow
                    :label="$gettext('Subject')"
                    placeholder=""
                    :rows="1"
                    color="secondary-dark"
                    class="input-subject"
                ></v-textarea>
              </v-flex>

              <v-flex xs12 sm6 md3 class="pa-2">
                <v-text-field
                    v-model="model.Details.Artist"
                    :append-icon="model.Details.ArtistSrc === 'manual' ? 'check' : ''"
                    :disabled="disabled"
                    :rules="[textRule]"
                    hide-details
                    browser-autocomplete="off"
                    :label="$gettext('Artist')"
                    placeholder=""
                    color="secondary-dark"
                    class="input-artist"
                ></v-text-field>
              </v-flex>

              <v-flex xs12 sm6 md3 class="pa-2">
                <v-text-field
                    v-model="model.Details.Copyright"
                    :append-icon="model.Details.CopyrightSrc === 'manual' ? 'check' : ''"
                    :disabled="disabled"
                    :rules="[textRule]"
                    hide-details
                    browser-autocomplete="off"
                    :label="$gettext('Copyright')"
                    placeholder=""
                    color="secondary-dark"
                    class="input-copyright"
                ></v-text-field>
              </v-flex>

              <v-flex xs12 sm6 md3 class="pa-2">
                <v-textarea
                    v-model="model.Details.License"
                    :append-icon="model.Details.LicenseSrc === 'manual' ? 'check' : ''"
                    :disabled="disabled"
                    :rules="[textRule]"
                    hide-details
                    browser-autocomplete="off"
                    auto-grow
                    :label="$gettext('License')"
                    placeholder=""
                    :rows="1"
                    color="secondary-dark"
                    class="input-license"
                ></v-textarea>
              </v-flex>

              <v-flex xs12 class="pa-2">
                <v-textarea
                    v-model="model.Description"
                    :append-icon="model.DescriptionSrc === 'manual' ? 'check' : ''"
                    :disabled="disabled"
                    hide-details
                    browser-autocomplete="off"
                    auto-grow
                    :label="$gettext('Description')"
                    placeholder=""
                    :rows="1"
                    color="secondary-dark"
                    class="input-description"
                ></v-textarea>
              </v-flex>

              <v-flex xs12 md6 class="pa-2">
                <v-textarea
                    v-model="model.Details.Keywords"
                    :append-icon="model.Details.KeywordsSrc === 'manual' ? 'check' : ''"
                    :disabled="disabled"
                    hide-details
                    browser-autocomplete="off"
                    auto-grow
                    :label="$gettext('Keywords')"
                    placeholder=""
                    :rows="1"
                    color="secondary-dark"
                    class="input-keywords"
                ></v-textarea>
              </v-flex>

              <v-flex xs12 md6 class="pa-2">
                <v-textarea
                    v-model="model.Details.Notes"
                    :append-icon="model.Details.NotesSrc === 'manual' ? 'check' : ''"
                    :disabled="disabled"
                    hide-details
                    browser-autocomplete="off"
                    auto-grow
                    :label="$gettext('Notes')"
                    placeholder=""
                    :rows="1"
                    color="secondary-dark"
                    class="input-notes"
                ></v-textarea>
              </v-flex>

              <v-flex v-if="!disabled" xs12 :text-xs-right="!rtl" :text-xs-left="rtl" class="pt-3">
                <v-btn depressed color="secondary-light" class="action-close"
                       @click.stop="close">
                  <translate>Close</translate>
                </v-btn>
                <v-btn color="primary-button" depressed dark class="action-apply action-approve"
                       @click.stop="save(false)">
                  <span v-if="$config.feature('review') && model.Quality < 3"><translate>Approve</translate></span>
                  <span v-else><translate>Apply</translate></span>
                </v-btn>
                <v-btn color="primary-button" depressed dark class="action-done hidden-xs-only"
                       @click.stop="save(true)">
                  <translate>Done</translate>
                </v-btn>
              </v-flex>
            </v-layout>
          </v-flex>
        </v-layout>

        <div class="mt-5"></div>
      </v-form>
    </v-container>
  </div>
</template>

<script>
import countries from "options/countries.json";
import Thumb from "model/thumb";
import * as options from "options/options";

export default {
  name: 'PTabPhotoDetails',
  props: {
    model: Object,
    uid: String,
  },
  data() {
    return {
      disabled: !this.$config.feature("edit"),
      config: this.$config.values,
      all: {
        colors: [{label: this.$gettext("Unknown"), name: ""}],
      },
      readonly: this.$config.get("readonly"),
      options: options,
      countries: countries,
      showDatePicker: false,
      showTimePicker: false,
      invalidDate: false,
      time: "",
      textRule: v => v.length <= this.$config.get('clip') || this.$gettext("Text too long"),
      rtl: this.$rtl,
    };
  },
  computed: {
    cameraOptions() {
      return this.config.cameras;
    },
    lensOptions() {
      return this.config.lenses;
    },
  },
  watch: {
    model() {
      this.updateTime();
    },
    uid() {
      this.updateTime();
    },
  },
  created() {
    this.updateTime();
  },
  methods: {
    updateTime() {
      if (!this.model.hasId()) {
        return;
      }

      const taken = this.model.getDateTime();

      this.time = taken.toFormat("HH:mm:ss");
    },
    updateModel() {
      if (!this.model.hasId()) {
        return;
      }

      let localDate = this.model.localDate(this.time);

      this.invalidDate = !localDate.isValid;

      if (this.invalidDate) {
        return;
      }

      if (this.model.Day === 0) {
        this.model.Day = parseInt(localDate.toFormat("d"));
      }

      if (this.model.Month === 0) {
        this.model.Month = parseInt(localDate.toFormat("L"));
      }

      if (this.model.Year === 0) {
        this.model.Year = parseInt(localDate.toFormat("y"));
      }

      const isoTime = localDate.toISO({
        suppressMilliseconds: true,
        includeOffset: false,
      }) + "Z";

      this.model.TakenAtLocal = isoTime;

      if (this.model.currentTimeZoneUTC()) {
        this.model.TakenAt = isoTime;
      }
    },
    left() {
      this.$emit('next');
    },
    right() {
      this.$emit('prev');
    },
    openPhoto() {
      this.$viewer.show(Thumb.fromFiles([this.model]), 0);
    },
    save(close) {
      if (this.invalidDate) {
        this.$notify.error(this.$gettext("Invalid date"));
        return;
      }

      this.updateModel();

      this.model.update().then(() => {
        if (close) {
          this.$emit('close');
        }

        this.updateTime();
      });
    },
    close() {
      this.$emit('close');
    },
  },
};
</script>
